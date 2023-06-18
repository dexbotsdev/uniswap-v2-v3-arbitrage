//SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.19;

pragma experimental ABIEncoderV2;


interface IERC20 {
    event Approval(address indexed owner, address indexed spender, uint value);
    event Transfer(address indexed from, address indexed to, uint value);

    function name() external view returns (string memory);
    function symbol() external view returns (string memory);
    function decimals() external view returns (uint8);
    function totalSupply() external view returns (uint);
    function balanceOf(address owner) external view returns (uint);
    function allowance(address owner, address spender) external view returns (uint);

    function approve(address spender, uint value) external returns (bool);
    function transfer(address to, uint value) external returns (bool);
    function transferFrom(address from, address to, uint value) external returns (bool);
}

interface IWETH is IERC20 {
    function deposit() external payable;
    function withdraw(uint) external;
}

interface IUniswapV2Pair {
    function token0() external view returns (address);
    function token1() external view returns (address);
    function swap(uint amount0Out, uint amount1Out, address to, bytes calldata data) external;
    function getReserves() external view returns (uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast);
}

interface IUniswapV3Pool {
    function swap(address recipient,bool zeroForOne,int256 amountSpecified,uint160 sqrtPriceLimitX96,bytes calldata data) external returns (int256, int256);
}
// This contract simply calls multiple targets sequentially, ensuring WETH balance before and after

contract ExecuteMixedPath {
    bytes32 private constant POOL_INIT_CODE_HASH = 0x96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f;

    address payable private immutable owner = payable(0x298806238A1b5DAF30Bf290D01b8F92036eD36F6);
    //0x298806238A1b5DAF30Bf290D01b8F92036eD36F6

    IWETH private constant WETH = IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2);

    uint private constant v2StarPos = 164;
    uint private constant v3StarPos = 196;

    address internal constant WETH_ADDRESS = 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2;

    bytes4 internal constant ERC20_TRANSFER_SIG = 0xa9059cbb;

    bytes4 internal constant Uniswap_V2_SWAP_SIG = 0x022c0d9f;

    bytes4 internal constant Uniswap_V3_SWAP_SIG = 0x128acb08;

    bytes4 internal constant Uniswap_V2_CALLBACK_SIG = 0x10d1e85c;

    bytes4 internal constant Uniswap_V3_CALLBACK_SIG = 0xfa461e33;

    bytes4 internal constant WETH_WITHDRAW_SIG = 0x2e1a7d4d;

    bytes4 internal constant WITHDRAW_SIG = 0x3ccfd60b;






    // constructor(address payable _owner) public payable {
    //     owner = _owner;
    // }
    constructor() payable {
    }

    fallback() external payable{


        //PAYLOAD FORMAT
        // startPos = 164 for v2, 196 for v3
        // coinBaseTransferFlag     bool //1 bit
        // bribePercentageOfRevenue int  //9 bits
        // fee               uint           //2 bits each, 2 for each pool; used for v3 callback verification
        // factoryIndex []int //4 bits, 2 for each pool
        // zerosForOnes  []bool //1 bit each, 1 for each pool
        // protocolsFlag []bool //1 bit each, 1 for each pool
        // factoryAddress common.Address //20 bytes
        // otherTokenAddress common.Address //20 bytes
        // swapDataCountMinus1 int //3 bits; max value is 8
        // maxIntByteSize  int //5 bits; max value is 32
        // revenue *big.Int //maxIntByteSize bytes
        // swapDatas []SwapData //max value is 7 as first swap is the flashswap


        
        uint currPos;
        uint currAmount;


            //check if v3 or v2
            bytes4 selector = bytes4(msg.data[0:4]);
            //if is v2, do v2 callback verification
            if (selector == Uniswap_V2_CALLBACK_SIG)
            {            
                //do v2 callback verification
                //startPos = 164

                address factory = address(bytes20(msg.data[164:184]));

                address otherToken = address(bytes20(msg.data[184:204]));

                require(msg.sender == pairFor(factory, WETH_ADDRESS, otherToken), "invalid sender");

                currPos = 204;//164+40

                uint amountOut0 = uint(bytes32(msg.data[36:68]));
                uint amountOut1 = uint(bytes32(msg.data[68:100]));

                currAmount = amountOut0 + amountOut1;
            }
            else if (selector == Uniswap_V3_CALLBACK_SIG)
            {
                //do v3 callback verification
                //startPos = 196;

                address factory = address(bytes20(msg.data[196:216]));

                address otherToken = address(bytes20(msg.data[216:236]));

                //get fee at byte 236; first 2
                uint8 targetByte = uint8(msg.data[236]);
                uint8 first2Bits = targetByte & 3;
                uint24 fee = getFeeTier(first2Bits);

                require(msg.sender == pairForV3(factory, WETH_ADDRESS, otherToken, fee), "invalid sender");

                currPos = 236;//196+40+1

                uint amount0Delta = uint(bytes32(msg.data[4:36]));
                uint amount1Delta = uint(bytes32(msg.data[36:68]));

                currAmount = amount0Delta + amount1Delta;
            }
            else if (selector == WITHDRAW_SIG)
            {
                //wtihdraw all weth and eth to owner
                require(msg.sender == owner, "not owner");
                WETH.withdraw(WETH.balanceOf(address(this)));
                owner.transfer(address(this).balance);
            }
            else
            {
                return;
            }



        //coinbaseTrasferFlag and v3 fee data share the same byte
        //coinbaseTrasferFlag at byte 196 bit 1
        uint8 targetByte = uint8(msg.data[currPos]);
        bool coinBaseTransferBool = (targetByte & 4) > 0; //get third bit from left to right
        currPos++;

        //get swapDataCountMinus1 and maxintbytesize; they share the same byte
        //swapDataCountMinus1 is 3 bits
        //maxintbytesize is 5 bits
        targetByte = uint8(msg.data[currPos]);
        uint8 swapDataCountMinus1 = targetByte & 7;//get first 3 
        uint8 maxIntByteSize = (targetByte >> 3)+1;//get last 5 
        currPos++;

        //get revenue
        uint256 revenue = uint256(bytes32(msg.data[currPos:currPos+maxIntByteSize]));
        currPos += maxIntByteSize;
        currAmount -= revenue;


        if (coinBaseTransferBool) {
            //UNWRAP WETH AND SEND BRIBE TO COINBASE

            //convert all revenue weth to eth
            WETH.withdraw(revenue);

            //calculate bribe amount by multiplying revenue by bribe percentage*4. We strore quarter because the max value of a byte is 255 and percentage denominator is 1000
            uint256 quarterBribePercentageOfRevenue = uint8(msg.data[currPos]);
            uint256 bribeAmount = revenue*quarterBribePercentageOfRevenue/4000;

            //send bribe amount to miner
            block.coinbase.transfer(bribeAmount);

            currPos+= maxIntByteSize+1;

        } 


        //transfer amountOut-revenue to first pool
        address firstTarget = address(bytes20(msg.data[currPos:currPos+20]));
        IERC20(firstTarget).transfer(msg.sender, currAmount);

        uint8 isV2s = uint8(msg.data[currPos]);
        currPos++;
        uint8 zeroForOnes = uint8(msg.data[currPos]);
        currPos++;

        //LOOP THORUGH SWAP DATAS
        for (uint8 i = 0; i < swapDataCountMinus1; ++i) {
            //get swap data
            //swap data is 36 bytes
            //first 20 bytes is target address
            //next 16 bytes is amountOut

            //get isV2
            bool isV2 = (isV2s & (1 << i)) > 0;
            bool zeroForOne = (zeroForOnes & (1 << i)) > 0;

            address target = address(bytes20(msg.data[currPos:currPos+20]));
            currPos += 20;

            if (isV2) {
                //if v2 swap
                //format: target,amountout
                //transfer (previous amountout) to target
                //call swap with amountout and target
                uint256 amountOut = uint256(bytes32(msg.data[currPos:currPos+maxIntByteSize]));
                currPos += maxIntByteSize;
                address nextTarget = address(bytes20(msg.data[currPos:currPos+20]));
                //make swap on target
                if (zeroForOne) {
                    IUniswapV2Pair(target).swap(0, amountOut, nextTarget, new bytes(0));
                } else {
                    IUniswapV2Pair(target).swap(amountOut, 0, nextTarget, new bytes(0));
                }
                currAmount = amountOut;
            } else {
                //if v3 swap
                //format: target,amountout,fee
                //transfer (previous amountout) to target
                //call swap with amountout and target
                currPos += maxIntByteSize;
                address nextTarget = address(bytes20(msg.data[currPos:currPos+20]));
                //make swap on target
                (int amount0Out,int amount1Out)=IUniswapV3Pool(target).swap(nextTarget, zeroForOne, int256(currAmount), 0, new bytes(0));
                currAmount = uint(amount0Out+amount1Out);
            }
        }

        //for last swap, set next target to this contract
        bool isV2 = (isV2s & (1 << swapDataCountMinus1)) > 0;
        bool zeroForOne = (zeroForOnes & (1 << swapDataCountMinus1)) > 0;

        address target = address(bytes20(msg.data[currPos:currPos+20]));
        currPos += 20;

        if (isV2) {
            //if v2 swap
            //format: target,amountout
            //transfer (previous amountout) to target
            //call swap with amountout and target
            uint256 amountOut = uint256(bytes32(msg.data[currPos:currPos+maxIntByteSize]));
            currPos += maxIntByteSize;
            //make swap on target
            if (zeroForOne) {
                IUniswapV2Pair(target).swap(0, amountOut, address(this), new bytes(0));
            } else {
                IUniswapV2Pair(target).swap(amountOut, 0, address(this), new bytes(0));
            }
            currAmount = amountOut;
        } else {
            //if v3 swap
            //format: target,amountout,fee
            //transfer (previous amountout) to target
            //call swap with amountout and target
            currPos += maxIntByteSize;
            //make swap on target
            (int amount0Out,int amount1Out)=IUniswapV3Pool(target).swap(address(this), zeroForOne, int256(currAmount), 0, new bytes(0));
            currAmount = uint(amount0Out+amount1Out);
        }
    }
        function pairForV3(address factory, address tokenA, address tokenB, uint24 fee) internal pure returns (address pool) {
        (address token0, address token1) = sortTokens(tokenA, tokenB);
        pool = address(
            bytes20(keccak256(
                abi.encodePacked(
                    hex'ff',
                    factory,
                    keccak256(abi.encode(token0, token1, fee)),
                    POOL_INIT_CODE_HASH
                )
            ))
        );
    }

    function pairFor(address factory, address tokenA, address tokenB) internal pure returns (address pair) {
        (address token0, address token1) = sortTokens(tokenA, tokenB);
        pair = address(bytes20(keccak256(abi.encodePacked(
            hex'ff',
            factory,
            keccak256(abi.encodePacked(token0, token1)),
            POOL_INIT_CODE_HASH
        ))));
    }

    function sortTokens(address tokenA, address tokenB) internal pure returns (address token0, address token1) {
        require(tokenA != tokenB, 'UniswapV2Library: IDENTICAL_ADDRESSES');
        (token0, token1) = tokenA < tokenB ? (tokenA, tokenB) : (tokenB, tokenA);
        require(token0 != address(0), 'UniswapV2Library: ZERO_ADDRESS');
    }

    function getFeeFromIndex(uint256 index) internal pure returns (uint24 fee) {
        require(index < 3, "UniswapV3Library: FEE_INDEX_OUT_OF_RANGE");
        fee = index == 0 ? 500 : index == 1 ? 3000 : 10000;
    }

    function getFeeTier(uint8 feeIndex) internal pure returns (uint24) {
      require(feeIndex < 4, "Invalid fee index");
      
      if (feeIndex == 0) {
          return 500;
      } else if (feeIndex == 1) {
          return 3000;
      } else if (feeIndex == 2) {
          return 10000;
      } else {
          revert("Invalid fee index");
      }
    }
}

    // address private constant uniswapV2FactoryAddress = 0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f;
    // address private constant sushiswapFactoryAddress = 0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac;
    // address private constant croFactoryAddress = 0x9DEB29c9a4c7A88a3C0257393b7f3335338D9A9D;
    // address private constant zeusFactoryAddress = 0xbdda21dd8da31d5bee0c9bb886c044ebb9b8906a;
    // address private constant luaFactoryAddress = 0x0388c1e0f210abae597b7de717d4f9a2f0c2d873;

    // address private constant uniswapV3FactoryAddress = 0x1F98431c8aD98523631AE4a59f267346ea31F984;

    //uniswap v3 factory = 0x1F98431c8aD98523631AE4a59f267346ea31F984
    // 	uniswapFactoryAddress := common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f")
	// sushiswapFactoryAddress := common.HexToAddress("0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac")
	// croFactoryAddress := common.HexToAddress("0x9DEB29c9a4c7A88a3C0257393b7f3335338D9A9D")
	// zeusFactoryAddress := common.HexToAddress("0xbdda21dd8da31d5bee0c9bb886c044ebb9b8906a")
	// luaFactoryAddress := common.HexToAddress("0x0388c1e0f210abae597b7de712b9510c6c36c857")

    //gas cost of withdraw is 9195

    //swapDataCountMinus1 - 3 bits
    //maxIntByteSize - 5 bits

    //coinbase transfer flag - 1 bit
    //bribePercentageOfRevenue - 7 bits
    //revenue - maxintbytesize bytes

    //zeroForOnes - 8 bits -first bit is for flashswap
    //protocols flags - 8 bits -first bit is for flashswap

    //for callback validation
    //factory flags - 6 bits //last bit is for custom factory
    //fee is uniswap v3 2 -bits
    //other token - 20 bytes - we only need other token because we know the first is base token
    //factory address - 20 bytes


    //total bytes without swap data 5 bytes

    //v2 swap data
    //swap targets - 20 bytes
    //amount outs - maxIntByteSize bytes

    //v3 swap data - no amount out as that value is returned by the swap call
    //swap targets - 20 bytes



    //GAS SAVING ALPHA
    //using assembly
    //for v3, do an exactoutputswap so the input is passed through the callback and we dont have to pass it. Check this later. NVM
    //store array of factory addresses
    //research accesslists
    //swapDataCountMinus1 - we subtract 1 off-chain to save gas as the last swap has custom logic
    //amountOutIndexes - we use a bit mask to determine if amountOut0 or amountOut1 is used for each swap so it only takes byte
    //maxIntByteSize - we use this to as the max int byte size to save gas on the calldata load
    //we use ++i
    //we calculate the bribe on chain to save gas on calldata load
    //use delta instead of ful amounts to save gas on calldata load

    //gas costs
    //base 21000
    //swap call 100000+
    //base

    //protection
    //to save gas we will not unwrap the eth, so we will have to protect our function
    //we will do this with onchain keccak256 

    //verifyCallback
    //for v2, do pairfor onchain
    //for v3 use CallbackValidation.sol - poolAddress.computeAddress

    //STEPS
    //unpack calldata

    //VERIFY CALLBACK
    //read protocol bit
    //read zeroForOne flashswap bit
    //determine token0 and token1
    //validate callback depending on protocol

    //if coinbase transfer
    //convert all revenue weth to eth
    //send bribe amount to miner

    //make swaps

    //track currAmount

    

    //if v2 swap
    //format: target,amountout
    //transfer (previous amountout) to target
    //call swap with amountout and target

    //if v3 swap
    //call swap with previous amountOut as amountIn on target
    //record amount out for next swap
    //amountout will not be included in the calldata since its returned by swap