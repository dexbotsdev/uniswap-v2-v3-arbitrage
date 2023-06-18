//SPDX-License-Identifier: UNLICENSED
pragma solidity 0.6.12;

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

struct SwapData{
    address target;
    uint256 amount1Out;
    uint256 amount0Out;
}

contract FBFSV3 {
    address internal constant owner = 0x298806238A1b5DAF30Bf290D01b8F92036eD36F6;

    address internal constant WETH_ADDRESS = 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2;

    bytes4 internal constant ERC20_TRANSFER_SIG = 0xa9059cbb;

    bytes4 internal constant PAIR_SWAP_SIG = 0x022c0d9f;

    bytes4 internal constant WETH_WITHDRAW_SIG = 0x2e1a7d4d;

    constructor() public payable {
    }

    receive() external payable {
    }

    function uniswapV2Call(address _sender, uint _amountOut0, uint _amountOut1, bytes calldata _data) external
    {
        //(uint256 revenue, SwapData[] memory swapData) = abi.decode(_data, (uint256, SwapData[]));

        
        
        //transfer profits to owner



        assembly {
            //Steps
            //1. Do Flashswap.. now we are in the callback function
            //2. withdraw WETH to ETH
            //3. transfer ETH revenue amount to owner
            //4. transfer remaining WETH to first target

            //Calldata format
            //0-3: function selector
            //4-35: sender address
            //36-67: amount0
            //68-99: amount1
            //***callback payload begins***
            //100-131: revenue
            //132-163: maxSwapIndex - 1 byte/ allows up to ~32 hops
            //164-195: amountOutIndexes - 1 byte / 8 bits - each bit represents if amountOut is 0 or 1 in order. Allows up to 8 hops in addition to the flashswap hop
            //196+: individual swap data, each 52 bytes long
            //(target,amouuntIn) - (addresses - 20 bytes each, uint256 - 32 bytes each) total 52 bytes




            //let amount0 := calldataload(36) 
            //let amount1 := calldataload(68)
            //let revenue := calldataload(164)
            //let maxSwapIndex := shr(248, calldataload(196))
            //let amountOutIndexes := shr(248, calldataload(197)) 
            //start if bytes is 36
            let amountOut0 := calldataload(36)//amount out 0
            let amountOut1 := calldataload(68)//amount out 1
            let revenue := calldataload(164)
            let maxSwapIndex := shr(248, calldataload(196))             //swap amounts
            let amountOutIndexes := shr(248, calldataload(197))            // bytes1 each bit represents if amountOut is 0 or 1 in order


            //WETH.withdraw(revenue);
            mstore(124, WETH_WITHDRAW_SIG)
            mstore(128, revenue)
            let s1 := call(gas(), WETH_ADDRESS, 0, 124, 36, 0, 0)


            //owner.transfer(revenue); 
            let s2 := call(gas(), owner, revenue, 0, 0, 0, 0)

            let firstTarget := shr(96, calldataload(198))

            //WETH transfer to first target by calculating the amount of WETH left after the swap
            //WETH.transfer(swapData[0].target, _amount0+_amount1-revenue); //we add amount0 and amount1 to get amountout regardless of index
            mstore(124, ERC20_TRANSFER_SIG)
            mstore(128, firstTarget)// destination
            mstore(160, sub(add(amountOut0,amountOut1),revenue))// amount
            let s3 := call(gas(),WETH_ADDRESS,0,124,68,0,0)

            //multi call the rest of the path
            //for (uint256 i = 0; i < swapData.length-2; i++) 
                //IUniswapV2Pair(swapData[i].target).swap(swapData[i].amount0Out,swapData[i].amount1Out,swapData[i+1].target,"");

            //calculate start position of the individual swap data
            currPos := 198
            swapDataSize := 52

            for { let i := 0 } lt(i, maxSwapIndex) { i := add(i, 1) } 
            {  
                //mask checks if bit i is 0 or 1
                let mask := shl(i,1)
                let amountOutIndex := iszero(and(mask,amountOutIndexes))

                let target := shr(96, calldataload(currPos))//20 bytes
                let amountOut := calldataload(add(currPos,20))
                let to := shr(96, calldataload(add(currPos,swapDataSize)))

                
                mstore(0x7c, PAIR_SWAP_SIG)// swap function signature
                switch amountOutIndex// amountOutIndex == 0 ? ....
                case 0 {
                    mstore(0x80, 0)
                    mstore(0xa0, amountOut)
                }
                default {
                    mstore(0x80, amountOut)
                    mstore(0xa0, 0)
                }
                mstore(0xc0, to)// address(this)
                mstore(0xe0, 0x80)// empty bytes


                let s := call(gas(), target, 0, 0x7c, 0xa4, 0, 0)

                //increment position by i*52. 52 is the size of the data for each swap
                currPos := add(currPos, mul(i,swapDataSize))

            }

            let mask := shl(maxSwapIndex,1)
            let amountOutIndex := iszero(and(mask,amountOutIndexes))

            let currPos := add(198, mul(maxSwapIndex,swapDataSize))

            let target := shr(96, calldataload(currPos))//20 bytes
            let amountOut := calldataload(add(currPos,20))
            //let to := shr(96, calldataload(add(currPos,52)))

                // swap function signature
            mstore(0x7c, PAIR_SWAP_SIG)
                // tokenOutNo == 0 ? ....
            switch amountOutIndex
            case 0 {
                mstore(0x80, 0)
                mstore(0xa0, amountOut)
            }
            default {
                mstore(0x80, amountOut)
                mstore(0xa0, 0)
            }
                // address(this)
            mstore(0xc0, caller())
                // empty bytes
            mstore(0xe0, 0x80)

            let s4 := call(gas(), target, 0, 0x7c, 0xa4, 0, 0)


            //weth withdraw
            //

        }



        //multi call the rest of the path
        //for (uint256 i = 0; i < swapData.length-2; i++) 
        //{
            //IUniswapV2Pair(swapData[i].target).swap(swapData[i].amount0Out,swapData[i].amount1Out,swapData[i+1].target,"");
        ///}

        //last call
        //IUniswapV2Pair(swapData[swapData.length-1].target).swap(swapData[swapData.length-1].amount0Out,swapData[swapData.length-1].amount1Out,msg.sender,"");
    }
}