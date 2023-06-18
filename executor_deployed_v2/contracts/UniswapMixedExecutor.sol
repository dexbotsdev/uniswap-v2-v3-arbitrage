//SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.19;

pragma experimental ABIEncoderV2;

//import "forge-std/Test.sol";


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

interface IUniswapV3Pool {
    function swap(address recipient,bool zeroForOne,int256 amountSpecified,uint160 sqrtPriceLimitX96,bytes calldata data) external returns (int256, int256);
    function token0() external view returns (address);
    function token1() external view returns (address);
    function fee() external view returns (uint24);
}

interface IUniswapV2Pair {
    function token0() external view returns (address);
    function token1() external view returns (address);
    function swap(uint amount0Out, uint amount1Out, address to, bytes calldata data) external;
    function getReserves() external view returns (uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast);
}

interface IUniswapV3Quoter {
    function quoteExactInputSingle(
        address tokenIn,
        address tokenOut,
        uint24 fee,
        uint256 amountIn,
        uint160 sqrtPriceLimitX96
    ) external returns (uint256 amountOut);
}


// This contract simply calls multiple targets sequentially, ensuring WETH balance before and after

contract UniswapMixedExecutor {
    bytes32 private constant POOL_INIT_CODE_HASH = 0x96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f;

    address payable private immutable owner = payable(0x71296ebC93BB8645Fc0826EAED445e55b0813B41);
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

    /// @dev The minimum value that can be returned from #getSqrtRatioAtTick. Equivalent to getSqrtRatioAtTick(MIN_TICK)
    uint160 internal constant MIN_SQRT_RATIO_PLUS_1 = 4295128740;
    /// @dev The maximum value that can be returned from #getSqrtRatioAtTick. Equivalent to getSqrtRatioAtTick(MAX_TICK)
    uint160 internal constant MAX_SQRT_RATIO_MINUS_1 = 1461446703485210103287273052203988822378723970341;

    // constructor(address payable _owner) public payable {
    //     owner = _owner;
    // }
    constructor() payable {
    }

    fallback() external payable{
        //// console.log("***FALLBACK***FALLBACK***FALLBACK***FALLBACK***FALLBACK***FALLBACK***FALLBACK***FALLBACK***FALLBACK***FALLBACK***");
        //// console.log("***FALLBACK***FALLBACK***FALLBACK***FALLBACK***FALLBACK***FALLBACK***FALLBACK***FALLBACK***FALLBACK***FALLBACK***");
    
        if (msg.data.length ==0){
            //fallback payable
            return;
        }


        //check if v3 or v2
        bytes4 selector = bytes4(msg.data[0:4]);
        // console.log("selector:" );
        // console.logBytes4(selector);

        uint mainPayloadStartPos;
        bool v3SwapsDone = false;

        //uint tokensLength;
        
        //if is v2, do v2 callback verification
        if (selector == Uniswap_V2_CALLBACK_SIG){            
            // console.log("***V2 CALLBACK***V2 CALLBACK***V2 CALLBACK***V2 CALLBACK***V2 CALLBACK***V2 CALLBACK***V2 CALLBACK***V2 CALLBACK***V2 CALLBACK***V2 CALLBACK***");
            // console.log("***V2 CALLBACK***V2 CALLBACK***V2 CALLBACK***V2 CALLBACK***V2 CALLBACK***V2 CALLBACK***V2 CALLBACK***V2 CALLBACK***V2 CALLBACK***V2 CALLBACK***");
            //do v2 callback verification
            //mainPayloadStartPos = 164

            //unpack V2CallbackV3SwapsDoneAndFactoryIndexAndOtherTokenIndexBytes []byte //1;1;3;3 bits
            uint8 targetByte  = uint8(msg.data[164]);
            bool v2CallbackV3SwapsDone = targetByte & 64 == 64;
            uint v2FactoryIndex = (targetByte & 56)>>3;
            uint v2OtherTokenIndex = targetByte & 7;


            //unpack tokenArr at index V2OtherTokenIndex

            uint currPos = 165;

            if (v2CallbackV3SwapsDone){
                v3SwapsDone = true;
            }

            address factory;
            if (v2FactoryIndex == 1){ //1 means custom address
                factory = address(bytes20(msg.data[165:185]));
                currPos += 20;
            } else {
                //select address from list
                factory = getV2Factory(v2FactoryIndex);
            }
            // console.log("facotry", factory);

            //unpack intByteSize
            uint8 intByteSize = uint8(msg.data[currPos]);
            //uint intRightShift = 256 - intByteSize*8;
            // console.log("intByteSize: %s", intByteSize);


            //unpack tokenArr at index V2OtherTokenIndex; pos = 4 + IntByteSize + V2OtherTokenIndex * 20
            uint tokensPos = currPos + 4 + intByteSize;
            address otherToken = address(bytes20(msg.data[tokensPos + v2OtherTokenIndex * 20: tokensPos + v2OtherTokenIndex * 20 + 20]));

            //sort weth and other token
            address token0 = WETH_ADDRESS;
            address token1 = otherToken;
            if (token0 > token1) (token0, token1) = (token1, token0);


            //verify v2 callback
            require(uniswapV2CallbackVerification(msg.sender, factory,token0,token1), "invalid sender");
        
            mainPayloadStartPos = currPos;
            //on to v3 swaps
        } else if (selector == Uniswap_V3_CALLBACK_SIG){//***V3 CALLBACK***
            // console.log("***V3 CALLBACK******V3 CALLBACK******V3 CALLBACK******V3 CALLBACK******V3 CALLBACK******V3 CALLBACK***V3 CALLBACK***V3 CALLBACK***V3 CALLBACK***V3 CALLBACK***");
            // console.log("***V3 CALLBACK******V3 CALLBACK******V3 CALLBACK******V3 CALLBACK******V3 CALLBACK******V3 CALLBACK***V3 CALLBACK***V3 CALLBACK***V3 CALLBACK***V3 CALLBACK***");
            //startPos = 132


            //unpack IsFlashswapV3SwapsDoneAndFeeIndexAndFactoryIndexBytes []byte //1;1;2;3 bits
            uint8 targetByte = uint8(msg.data[133]);
            //bool callbackIsFlashswap = targetByte & 64 == 64; //1st bit
            bool callbackV3SwapsDone = targetByte & 32 == 32; //5th bit
            uint8 feeIndex = (targetByte & 24) >> 3; //next 2 bits
            uint8 factoryIndex = targetByte & 7; //last 3 bits

            //unpack Token0IndexAndToken1IndexBytes []byte //3;3 bits
            targetByte = uint8(msg.data[134]);
            uint8 token0Index = targetByte >> 3;
            uint8 token1Index = targetByte & 7;

            if (callbackV3SwapsDone){
                v3SwapsDone = true;
            }
            
            uint24 fee = getFee(feeIndex);
            // console.log("fee: %s", fee);

            uint currPos = 135;

            //if factory index = 1, unpack factory address
            address factory;
            if (factoryIndex == 1){ //factory index = 1 means custom address
                factory = address(bytes20(msg.data[currPos:currPos+20]));
                currPos += 20;
            } else {
                //select address from list
                factory = getV3Factory(factoryIndex);
            }

            //unpack intByteSize
            uint8 intByteSize = uint8(msg.data[currPos]);
            // console.log("currPos: %s", currPos);
            // console.log("intByteSizeByte: %s", uint8(msg.data[currPos]));

            //unpack tokensLength from CoinbaseTransferAndV2SwapDataLengthAndTokensLengthBytes []byte //1;3;3 bits
            targetByte = uint8(msg.data[currPos+2]);
            uint8 tokensLength = uint8(msg.data[currPos+2]) & 7;
            // console.log("targetByte: %s", targetByte);
            // console.log("tokensLength: %s", tokensLength);

            uint tokensPos = currPos + 4 + intByteSize;
            // console.log("tokensPos: %s", tokensPos);

            //if isToken0Weth, set to weth, else unpack token0 address
            address token0;
            if (token0Index == tokensLength){//if token0Index == tokensLength, set to weth
                // console.log("token0Index", token0Index, "tokensLength", tokensLength);
                // console.log("token0Index == tokensLength");
                token0 = WETH_ADDRESS;
            } else {
                token0 = address(bytes20(msg.data[tokensPos + token0Index * 20: tokensPos + token0Index * 20 + 20]));
            }

            //if isToken1Weth, set to weth, else unpack token1 address
            address token1;
            if (token1Index == tokensLength){
                // console.log("token1Index", token1Index, "tokensLength", tokensLength);
                // console.log("token1Index == tokensLength");
                token1 = WETH_ADDRESS;
            } else {
                token1 = address(bytes20(msg.data[tokensPos + token1Index * 20: tokensPos + token1Index * 20 + 20]));
            }

            // console.log("factory: %s", factory);
            // console.log("token0: %s", token0);
            // console.log("token1: %s", token1);


            //verify callback
            require(uniswapV3CallbackVerification(msg.sender, factory, token0, token1, fee), "invalid sender");

            //set firstAmountIn
            //if SwapsIndex == 0, set amountIn to amount

            // console.log("***V3 CALLBACK VERIFIED***");


            mainPayloadStartPos = currPos;
        } else if (selector == WITHDRAW_SIG) {
            // console.log("***WITHDRAW***");
            //wtihdraw all weth and eth to owner
            //require(msg.sender == owner, "not owner");
            uint256 wethBalance = WETH.balanceOf(address(this));
            if (wethBalance > 0){
                WETH.withdraw(wethBalance);
            }
            owner.transfer(address(this).balance);
            return;
        } else if (msg.data.length > 0){//for generic call
            // console.log("***V3 INITIAL CALL***");
            mainPayloadStartPos = 0;
        } else {
            // console.log("***NO DATA***");
            return;
        }

        //set currPos to correct position

        //unpack IntByteSizeBytes []byte //1;6 bits
        uint intByteSize = uint8(msg.data[mainPayloadStartPos]);
        uint intRightShift = 256 - intByteSize*8;
        // console.log("intByteSize: %s", intByteSize);
        // console.log("intRightShift: %s", intRightShift);


        //unpack swapsLengthAndAmountsLengthBytes []byte //4;4 bits
        uint8 targetByte = uint8(msg.data[mainPayloadStartPos+1]);
        uint8 swapsLength = targetByte >> 4; //first 4 bits
        uint8 amountsLength = targetByte & 15; //last 4 bits
        // console.log("targetByte: %s", targetByte);
        // console.log("swapsLength: %s", swapsLength);
        // console.log("amountsLength: %s", amountsLength);

        //unpack tokens length from CoinbaseTransferAndV2SwapDataLengthAndTokensLengthBytes []byte //1;3;3 bits
        targetByte = uint8(msg.data[mainPayloadStartPos+2]);
        uint8 tokensLength = targetByte & 7; //last 3 bits

        //set start of arrays
        uint tokensStartPos = mainPayloadStartPos + 4 + intByteSize;
        uint targetsStartPos = tokensStartPos + tokensLength*20;
        uint amountStartPos = targetsStartPos + swapsLength*20;
        uint zeroForOnesPos = amountStartPos + amountsLength*intByteSize;
        // console.log("targetsStartPos: %s", targetsStartPos);
        // console.log("amountStartPos: %s", amountStartPos);
        // console.log("zeroForOnesPos: %s", zeroForOnesPos);

        uint8 zeroForOnes = uint8(msg.data[zeroForOnesPos]);
        // console.log("zeroForOnes: %s", zeroForOnes);

        uint currPos = zeroForOnesPos+1;
        if (!v3SwapsDone){
            // console.log("***START V3 SWAP***");
            
            //calculate v3 swap data position; pos = 4 + IntByteSize + swapsLength*20 + AmountsLength*IntByteSize + 1 
            uint v3SwapDataStartPos = currPos;
            // console.log("v3SwapDataStartPos: %s", v3SwapDataStartPos);

            //unpack V3SwapDataBytes struct
            // type V3SwapDataBytes struct {
            // 	IsLastV3SwapAndSwapIndexAndAmountIndexBytes []byte //1;3;3 bits
            // 	IsToken0WethAndIsToken1WethAndFeeIndexAndFactoryIndexBytes []byte //1;1;2;3 bits
            // 	FactoryAddressBytes []byte
            // 	Token0Bytes         []byte
            // 	Token1Bytes         []byte
            // }

            //unpack IsLastV3SwapAndSwapIndexAndAmountIndexBytes []byte //1;3;3 bits
            uint8 targetByte = uint8(msg.data[currPos]);   
            uint8 swapIndex = (targetByte & 56)>>3; //next 3 bits
            uint8 amountIndex = targetByte & 7; //last 3 bits
            currPos += 1;
            // console.log("targetByte: %s", targetByte);
            // console.log("swapIndex: %s", swapIndex);
            // console.log("amountIndex: %s", amountIndex);

            //unpack IsToken0WethAndIsToken1WethAndFeeIndexAndFactoryIndexBytes []byte //1;1;2;3 bits
            //fee is not unpacked as it is not needed
            targetByte = uint8(msg.data[currPos]);
            uint8 factoryIndex = targetByte & 7; //last 3 bits
            // console.log("targetByte: %s", targetByte);
            // console.log("factoryIndex: %s", factoryIndex);
            // console.log("*****************************************************************************************************");

            //unpack target and next target; pos = 4 + IntByteSize
            address target = address(bytes20(msg.data[targetsStartPos + swapIndex*20:targetsStartPos + (swapIndex+1)*20]));
            address nextTarget;
            if (swapIndex == swapsLength-1){//if last address, set next target to this contract
                nextTarget = address(this);
            } else {
                nextTarget = address(bytes20(msg.data[targetsStartPos + (swapIndex+1)*20:targetsStartPos + (swapIndex+2)*20]));
            }         

            //AmountsStartPos = mainPayloadStartPos + 4 + IntByteSize + swapsLength*20
            //unpack amount at amountIndex
            uint amount = uint(bytes32(msg.data[amountStartPos + amountIndex*intByteSize:amountStartPos + (amountIndex+1)*intByteSize])) >> intRightShift;   

            //unpackZeroForOne; pos = 4 + IntByteSize + swapsLength*20 + AmountsLength*IntByteSize
            uint mask = 1 << (swapIndex);
            bool zeroForOne = (zeroForOnes & mask) == mask;

            uint160 sqrtPriceLimitX96 = zeroForOne ? MIN_SQRT_RATIO_PLUS_1 : MAX_SQRT_RATIO_MINUS_1;

            //modify payload
            uint v3SwapDataEndPos = v3SwapDataStartPos + 3;
            if(factoryIndex == 1){// if factory is not indexed; value of 1 means custom factory address
                v3SwapDataEndPos += 20;
            }

            //move current v3 swap data to front of payload
            bytes memory currentV3SwapData = msg.data[v3SwapDataStartPos:v3SwapDataEndPos];
            // console.log("currentV3SwapData: ");
            // console.logBytes(currentV3SwapData);

            //if firstAmountIn > 0 then add firstAmountIn to and of payload
            //set to empty bytes 

            //for next payload we cutout this v3SwapData (cutout = v3SwapDataEndPos - v3SwapDataStartPos)
            bytes memory newPayloadHead = msg.data[mainPayloadStartPos:v3SwapDataStartPos];
            bytes memory newPayloadTail = msg.data[v3SwapDataEndPos:];
            bytes memory newPayload = abi.encodePacked(currentV3SwapData, newPayloadHead, newPayloadTail);


            // move callback verification section to the front of payload
            // bytes memory v3CallbackVerificationSection = msg.data[mainPayloadStartPos:mainPayloadStartPos+4+intByteSize+swapsLength*20];


            //make swap on target
            IUniswapV3Pool(target).swap(nextTarget, zeroForOne, int256(amount), sqrtPriceLimitX96, newPayload);
            // console.log("amount0Delta");
            // console.logInt(amount0Delta);
            // console.log("amount1Delta");
            // console.logInt(amount1Delta);

            // console.log("eth balance: %s", address(this).balance);
            return;
        } else {//v3 swaps finished
            // console.log("weth balance: %s", WETH.balanceOf(address(this)));
            // console.log("***V3 SWAPS FINISHED***");
            //unpack CoinbaseTransferAndV2SwapDataLengthAndTokensLengthBytes []byte //1;3;3 bits
            uint CoinbaseTransferAndV2SwapDataLengthAndTokensLengthBytesPos = mainPayloadStartPos + 2;
            uint8 targetByte = uint8(msg.data[CoinbaseTransferAndV2SwapDataLengthAndTokensLengthBytesPos]);
            bool coinbaseTransfer = targetByte & 64 == 64; //first bit
            uint8 v2SwapDataLength = (targetByte & 56)>>3; //next 3 bits

            // console.log("targetByte: %s", targetByte);
            // console.log("coinbaseTransfer: %s", coinbaseTransfer);
            // console.log("v2SwapDataLength: %s", v2SwapDataLength);


            //unpack revenue
            uint revenuePos = mainPayloadStartPos + 3;
            uint revenue = uint(bytes32(msg.data[revenuePos:revenuePos+intByteSize])) >> intRightShift;
            // console.log("revenue: %s", revenue);


            //unpack QuarterBribePercentageOfRevenueBytes []byte //8 bits; only if coinBaseTransferBool is true
            if (coinbaseTransfer){
                //coinbase transfer
                uint QuarterBribePercentageOfRevenueBytesPos = mainPayloadStartPos + 3 + intByteSize;
                uint8 quarterBribePercentageOfRevenue = uint8(msg.data[QuarterBribePercentageOfRevenueBytesPos]);
                // console.log("quarterBribePercentageOfRevenue", quarterBribePercentageOfRevenue);

                //calculate bribe amount
                uint bribeAmount = revenue * quarterBribePercentageOfRevenue / 250;
                // console.log("bribeAmount", bribeAmount);

                //withdraw eth revenue
                WETH.withdraw(revenue);

                //send bribe to coinbase
                block.coinbase.transfer(bribeAmount);
                // console.log("bribe sent to coinbase");
            }

            //transfer first amount to first target
            address firstTarget = address(bytes20(msg.data[targetsStartPos:targetsStartPos+20]));
            // console.log("firstTarget: %s", firstTarget);

            //get firstAmountIn at the first amount
            uint firstAmountIn = uint(bytes32(msg.data[amountStartPos:amountStartPos+intByteSize])>>intRightShift);
            // console.log("firstAmountIn: %s", firstAmountIn);


            bool success = WETH.transfer(firstTarget, firstAmountIn);
            require(success, "transfer to first target failed");

            // console.log("FIRST AMOUNT TRANSFERED");
            // console.log("START V2 SWAPS");

            //start v2 swaps
            // console.log("all V2 swap data: ");
            // console.logBytes(msg.data[currPos:]);

            // console.log("v2 swap data length: %s", v2SwapDataLength);
            
            //LOOP THORUGH V2 SWAP DATAS
            for (uint8 i = 0; i < v2SwapDataLength; ++i) {
                // console.log("SWAP:", i);

                // unpack SwapIndexAndAmountIndexBytes []byte //3;3 bits
                uint8 targetByte = uint8(msg.data[currPos]);
                uint swapIndex = targetByte >> 3; //first 3 bits
                uint amountIndex = targetByte & 7; //last 3 bits

                currPos ++;

                // console.log("swapIndex: %s", swapIndex);
                // console.log("amountIndex: %s", amountIndex);

                //get target, zeroForOne and amountOut
                address target = address(bytes20(msg.data[targetsStartPos + swapIndex*20:targetsStartPos + (swapIndex+1)*20]));
                address nextTarget = address(bytes20(msg.data[targetsStartPos + (swapIndex+1)*20:targetsStartPos + (swapIndex+2)*20]));
                uint amountOut = uint(bytes32(msg.data[amountStartPos + amountIndex*intByteSize:amountStartPos + (amountIndex+1)*intByteSize])) >> intRightShift;
                
                uint mask = 1 << swapIndex;
                bool zeroForOne = (zeroForOnes & mask) == mask;

                // console.log("target: %s", target);
                // console.log("nextTarget: %s", nextTarget);
                // console.log("amountOut: %s", amountOut);
                // console.log("zeroForOne: %s", zeroForOne);

                //print reserves
                //(uint reserve0, uint reserve1,) = IUniswapV2Pair(target).getReserves();
                // console.log("reserve0: %s", reserve0);
                // console.log("reserve1: %s", reserve1);

                // console.log("zeroForOnesByte: %s", uint8(msg.data[zeroForOnesPos]));

                //make swap on target
                if (zeroForOne) {
                    IUniswapV2Pair(target).swap(0, amountOut, nextTarget, new bytes(0));
                } else {
                    IUniswapV2Pair(target).swap(amountOut, 0, nextTarget, new bytes(0));
                }
            }

            // console.log("V2 SWAPS FINISHED");
            return; //end of swaps
        }
    }
    function uniswapV3CallbackVerification(address poolToBeVerified,address factory, address token0, address token1, uint24 fee) internal pure returns (bool) {
        // // console.log("***V3 CALLBACK VERIFICATION***");
        // // console.log("poolToBeVerified", poolToBeVerified);
        // // console.log("factory", factory);
        // // console.log("token0", token0);
        // // console.log("token1", token1);
        // // console.log("fee", fee);

        address pool = address(bytes20(uint160(uint256(keccak256(
                abi.encodePacked(
                    hex'ff',
                    factory,
                    keccak256(abi.encode(token0, token1, fee)),
                    hex'e34f199b19b2b4f47f68442619d555527d244f78a3297ea89325f843f87b8b54'
                )
            ))))
        );
        // console.log("keccak256result", pool);

        return pool == poolToBeVerified;
    }

    function uniswapV2CallbackVerification(address poolToBeVerified, address factory, address token0, address token1) internal pure returns (bool) {
        // console.log("***V2 CALLBACK VERIFICATION***");
        // console.log("poolToBeVerified", poolToBeVerified);
        // console.log("factory", factory);
        // console.log("token0", token0);
        // console.log("token1", token1);

        
        address pair = address(bytes20(uint160(uint256(keccak256(abi.encodePacked(
            hex'ff',
            factory,
            keccak256(abi.encodePacked(token0, token1)),
            hex'96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f'
        ))))));

        // console.log("keccak256result", pair);

        return pair == poolToBeVerified;
    }


    function getFee(uint8 feeIndex) internal pure returns (uint24) {
      require(feeIndex < 4, "Invalid fee index");
      
      if (feeIndex == 0) {
          return 500;
      } else if (feeIndex == 1) {
          return 3000;
      } else if (feeIndex == 2) {
          return 10000;
      } else if (feeIndex == 3) {
          return 100;
      } else {
          revert("Invalid fee index");
      }
    }
    
    function getV2Factory(uint index) internal pure returns (address){
        if (index ==0){
            //uniswap v2 factory
            return 0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f;
        }
        if (index == 2){
            //sushi v2 factory
            return 0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac;
        }
    }
    function getV3Factory(uint index) internal pure returns (address){
        if (index ==0){
            return 0x1F98431c8aD98523631AE4a59f267346ea31F984;
        }
    }
}

    //calldata segments /v3Swaps/SendAmountToFirstV2Market/v2Swaps/coinbaseTransfer/transferToOwner

    //calldata requirements
    //v3 calldata sections
    //1 byte bool to say if its the final V3 swap
    //recipient address

    //v3 segments
    //first byte bool to say if its the final V3 swap
    //decode amount deltas to get amount out
    //set next amountIn to amountOut
    //get recipient from calldata

    //coinbase and owner transfer section
    //bribe
    //revenue

    //v2 calldata sections
    //recipient address
    //amount out



    //layout

    //flashswap section(only if the flashswap pair is v2)

    //v3 section

    //v2 section

    //coinbase transfer section

    //HIGH LEVEL STEPS
    //flashswap last pool V2, Or V3
    //flashswap V3 pools in reverse order
    //do coinbase and owner transfers
    //do V2 swaps in order

    //CASES
    //V2 only - V2 flashswap
    //Mixed with V3 first - do V3s then V2s
    //Mixed with V2 first

    //MISC NOTES
    //last swap points to executor
    //for mixed paths the first path must be V3, leaving V2 paths will be added to the end
    //weth amount must be sent to first contract from executor
    //last pool must be flashswapped
    //transfers to owner and coinbase must be done before first swap if it is V2
    //for verification pass this value instead of calculating on chain keccak256(abi.encode(token0, token1, fee(v3 only))), fee is for V3 only


    //FUNCTION SELECTORS
    //v2 callback
    //v3 callback
    //initial call to start v3 swap
    //recieve payable(to recieve eth)
    //withdraw to owner(to withdraw eth)
    //generic target payload by owner(incase we need to call a target with a payload that is not a swap)

    //V2 callback
    // verfify callback


    //V3 callback
    //verify callback
    //determine if its the last v3 swap
    //if not last v3 swap call next v3 swap
    //if last v3 swap, proceed to coinbase/owner transfers

    //Initial call to start v3 swap
    //check owner
    //decode data and call swap

    //Transfer To Coinbase
    //check flag
    //withdraw all eth
    //send bribe amount of eth to coinbase

    //Transfer To Owner
    //check flag
    //withdraw all remaining eth

    //Do V2 Swaps




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
    //intByteSize - 5 bits

    //coinbase transfer flag - 1 bit
    //bribePercentageOfRevenue - 7 bits
    //revenue - intByteSize bytes

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
    //amount outs - intByteSize bytes

    //v3 swap data - no amount out as that value is returned by the swap call
    //swap targets - 20 bytes



    //GAS SAVING ALPHA
    //using assembly
    //for v3, do an exactoutputswap so the input is passed through the callback and we dont have to pass it. Check this later. NVM
    //store array of factory addresses
    //research accesslists
    //swapDataCountMinus1 - we subtract 1 off-chain to save gas as the last swap has custom logic
    //amountOutIndexes - we use a bit mask to determine if amountOut0 or amountOut1 is used for each swap so it only takes byte
    //intByteSize - we use this to as the max int byte size to save gas on the calldata load
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


