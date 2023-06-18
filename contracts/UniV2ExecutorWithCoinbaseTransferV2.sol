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
// This contract simply calls multiple targets sequentially, ensuring WETH balance before and after

contract ExecutorV1WithCoinbaseTransfer {
    address payable private immutable owner;
    IWETH private constant WETH = IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2);


    constructor(address payable _owner) public payable {
        owner = _owner;
    }

    receive() external payable {
    }

    function uniswapV2Call(address _sender, uint _amountOut0, uint _amountOut1, bytes calldata _data) external
    {
        //Steps
        //1. Do Flashswap.. now we are in the callback function
        //2. withdraw WETH to ETH
        //3. transfer ETH bribe to miner; we transfer bribe on chain so we can capture big arb amounts  when we dont have enough ETH in the EOA
        //4. transfer bribe to miner
        //5. transfer remaining ETH revenue amount to owner
        //6. transfer remaining WETH to first target
        //7. multi call the rest of the path


        //gas saving alpha
        //swapCountMinus1 - we subtract 1 off-chain to save gas as the last swap has custom logic
        //amountOutIndexes - we use a bit mask to determine if amountOut0 or amountOut1 is used for each swap so it only takes byte
        //maxIntByteSize - we use this to as the max int byte size to save gas on the calldata load
        //we use ++i
        //we calculate the bribe on chain to save gas on calldata load
        

        //Calldata format
        //0-3: function selector
        //4-35: sender address
        //36-67: amount0
        //68-99: amount1
        //***callback payload begins***
        //164-195: revenue - 32 bytes
        //196-227: bribe - 32 bytes
        //228-259: maxSwapIndex - 32 bytes
        //260: maxSwapIndex - 1 byte/ allows up to ~32 hops
        //261: amountOutIndexes - 1 byte / 8 bits - each bit represents if amountOut is 0 or 1 in order. Allows up to 8 hops in addition to the flashswap hop
        //262+i*52: individual swap data, each 52 bytes long



        //uint256 amountOut0; 
        //uint256 amountOut1;

        uint256 maxIntByteSize;

        uint256 revenue;

        uint256 swapCountMinus1;
        uint256 amountOutIndexes; 

        uint256 intRightShift;


        //array of {address, size, int}



        assembly{
            //amountOut0 := calldataload(36)//36-67
            //amountOut1 := calldataload(68)//68-99
            //payload data begins at 164
            //maxIntByteSize := shr(calldataload(164),248) //1 byte; byte size of the max int in the calldata
            //bribePercentageOfRevenue := shr(calldataload(165),248) //1 byte
            //swapCountMinus1 := shr(calldataload(166),248) //197 1 byte we subtract 1 off chain to save gas as the last swap has custom logic
            //amountOutIndexes := shr(calldataload(167),248) //1 byte
            //intRightShift := sub(256,mul(maxIntByteSize,8)) //8 bits per byte
            //revenue := shr(calldataload(168),intRightShift)//168-199

            maxIntByteSize := shr(calldataload(164),248) //1 byte; byte size of the max int in the calldata
            swapCountMinus1 := shr(calldataload(166),248) //197 1 byte we subtract 1 off chain to save gas as the last swap has custom logic
            amountOutIndexes := shr(calldataload(167),248) //1 byte

            intRightShift := sub(256,mul(maxIntByteSize,8)) //8 bits per byte

            revenue := shr(calldataload(168),intRightShift)//168-199
        }
        
        {
            uint256 bribePercentageOfRevenue; 

            assembly{
                bribePercentageOfRevenue := shr(calldataload(165),248) //1 byte
            }

            //withdraw WETH to ETH
            WETH.withdraw(revenue);

            //transfer ETH bribe to miner; we transfer bribe on chain so we can capture big arb amounts  when we dont have enough ETH in the EOA
            uint256 bribe = revenue * bribePercentageOfRevenue / 100;
            block.coinbase.transfer(bribe);

            //transfer remaining ETH revenue amount to owner
            owner.transfer(revenue-bribe);
        }
    }
    function _performSwaps(uint _amountOut0, uint _amountOut1, uint256 revenue, uint256 swapCountMinus1, uint256 amountOutIndexes, uint256 intRightShift, uint256 swapDataSize, uint256 currPos) internal {
        {
            uint256 bribePercentageOfRevenue; 

            assembly{
                bribePercentageOfRevenue := shr(calldataload(165),248) //1 byte
            }

            //withdraw WETH to ETH
            WETH.withdraw(revenue);

            //transfer ETH bribe to miner; we transfer bribe on chain so we can capture big arb amounts  when we dont have enough ETH in the EOA
            uint256 bribe = revenue * bribePercentageOfRevenue / 100;
            block.coinbase.transfer(bribe);

            //transfer remaining ETH revenue amount to owner
            owner.transfer(revenue-bribe);
        }


        uint256 currPos = 200;
        uint256 swapDataSize = maxIntByteSize + 20;


        uint256 mask = 1;
        bool amountOutIndex = !(mask & amountOutIndexes > 0);

        address target;
        uint256 amountOut;//+20 to skip the target address
        address to;

        assembly{
            target:=shr(calldataload(currPos),96)
            amountOut:=shr(calldataload(add(currPos,20)),intRightShift)
            to:= shr(calldataload(add(currPos,swapDataSize)),96)
        }

        //uint256 amountToFirstTarget = _amountOut0 + _amountOut1 - revenue;
        WETH.transfer(target,_amountOut0 + _amountOut1 - revenue);

        IUniswapV2Pair(target).swap(amountOutIndex ? amountOut : 0,amountOutIndex ? 0 : amountOut,to,"");


        for (uint256 i = 1; i < swapCountMinus1; ++i) {  
            //mask checks if bit i is 0 or 1
            mask = 1 << i;
            amountOutIndex = !(mask & amountOutIndexes > 0);

            assembly{
                target:=shr(calldataload(currPos),96)
                amountOut:=shr(calldataload(add(currPos,20)),intRightShift)
                to:= shr(calldataload(add(currPos,swapDataSize)),96)
            }

            IUniswapV2Pair(target).swap(amountOutIndex ? amountOut : 0,amountOutIndex ? 0 : amountOut,to,"");

            //increment position by i*52. 52 is the size of the data for each swap
            currPos += i * swapDataSize;
        }

        mask = 1 << swapCountMinus1;
        amountOutIndex = !(mask & amountOutIndexes > 0);

        assembly{
            target:=shr(calldataload(currPos),96)
            amountOut:=shr(calldataload(add(currPos,20)),intRightShift)
        }
        
        IUniswapV2Pair(target).swap(amountOutIndex ? amountOut : 0,amountOutIndex ? 0 : amountOut,msg.sender,"");
    }

}
