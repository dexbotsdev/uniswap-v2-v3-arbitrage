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

    function uniswapV2Call(address _sender, uint _amount0, uint _amount1, bytes calldata _data) external
    {
        //(uint256 revenue, SwapData[] memory swapData) = abi.decode(_data, (uint256, SwapData[]));

        
        
        //transfer profits to owner



        assembly {


            //start if bytes is 36
            let amount0 := calldataload(36)
            let amount1 := calldataload(68)
            let revenue := calldataload(164)
            let maxSwapIndex := shr(248, calldataload(196))
            let amountOutNos := shr(248, calldataload(197))            // bytes1


            //WETH.withdraw(revenue);
            mstore(124, WETH_WITHDRAW_SIG)
            mstore(128, revenue)
            let s1 := call(gas(), WETH_ADDRESS, 0, 124, 36, 0, 0)


            //owner.transfer(revenue); 
            let s2 := call(gas(), owner, revenue, 0, 0, 0, 0)

            let firstTarget := shr(96, calldataload(198))

            //WETH.transfer(swapData[0].target, _amount0+_amount-revenue);
            mstore(124, ERC20_TRANSFER_SIG)
            mstore(128, firstTarget)// destination
            mstore(160, sub(add(amount0,amount1),revenue))// amount
            let s3 := call(gas(),WETH_ADDRESS,0,124,68,0,0)



            //multi call the rest of the path
            //for (uint256 i = 0; i < swapData.length-2; i++) 
                //IUniswapV2Pair(swapData[i].target).swap(swapData[i].amount0Out,swapData[i].amount1Out,swapData[i+1].target,"");

            for { let i := 0 } lt(i, maxSwapIndex) { i := add(i, 1) } 
            {  
                let mask := shl(i,1)
                let amountOutNo := iszero(and(mask,amountOutNos))
                
                let startPos := add(198, mul(i,52))

                let target := shr(96, calldataload(startPos))//20 bytes
                let amountOut := calldataload(add(startPos,20))
                let to := shr(96, calldataload(add(startPos,52)))

                
                mstore(0x7c, PAIR_SWAP_SIG)// swap function signature
                switch amountOutNo// amountOutNo == 0 ? ....
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
            }

            let mask := shl(maxSwapIndex,1)
            let amountOutNo := iszero(and(mask,amountOutNos))

            let startPos := add(198, mul(maxSwapIndex,52))

            let target := shr(96, calldataload(startPos))//20 bytes
            let amountOut := calldataload(add(startPos,20))
            //let to := shr(96, calldataload(add(startPos,52)))

                // swap function signature
            mstore(0x7c, PAIR_SWAP_SIG)
                // tokenOutNo == 0 ? ....
            switch amountOutNo
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