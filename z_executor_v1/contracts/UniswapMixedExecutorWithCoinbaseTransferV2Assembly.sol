    //SPDX-License-Identifier: UNLICENSED
    pragma solidity 0.8.19;

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

        constructor() payable {
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

                //calldata retrival lines
                //maxIntByteSize := shr(248,calldataload(164)) //1 byte; byte size of the max int in the calldata
                //bribePercentageOfRevenue := shr(248,calldataload(165)) //1 byte
                //swapCountMinus1 := shr(248,calldataload(166)) //197 1 byte we subtract 1 off chain to save gas as the last swap has custom logic
                //amountOutIndexes := shr(248,calldataload(167)) //1 byte
                //revenue := shr(intRightShift,calldataload(168))//168-199


                let maxIntByteSize := shr(248,calldataload(164)) //1 byte; byte size of the max int in the calldata

                let intRightShift := sub(256,mul(maxIntByteSize,8)) //8 bits per byte

                // uint256 swapDataSize = maxIntByteSize + 20;
                let swapDataSize := add(maxIntByteSize,20)

                let currPos := 200 //200 is the start of the swap data
                
                let revenue := shr(intRightShift,calldataload(168))//168-199


                //avoid stack too deep
                {
                    let bribePercentageOfRevenue := shr(248,calldataload(165)) //1 byte


                    //WETH.withdraw(revenue);
                    mstore(124, WETH_WITHDRAW_SIG)
                    mstore(128, revenue)
                    pop(call(gas(), WETH_ADDRESS, 0, 124, 36, 0, 0))

                    //uint256 bribe = revenue * bribePercentageOfRevenue / 100;
                    let bribe := div(mul(revenue,bribePercentageOfRevenue),100)

                    //owner.transfer(revenue-bribe); 
                    pop(call(gas(), owner, sub(revenue,bribe), 0, 0, 0, 0))

                    //todo coinbase transfer
                    pop(call(gas(), coinbase(), bribe, 0, 0, 0, 0))
                }


                let target := shr(96, calldataload(currPos))//20 bytes
                let amountOut := shr(intRightShift,calldataload(add(currPos,20)))
                let to := shr(96, calldataload(add(currPos,swapDataSize)))

                //WETH transfer to first target by calculating the amount of WETH left after the swap
                //WETH.transfer(swapData[0].target, _amount0+_amount1-revenue); //we add amount0 and amount1 to get amountout regardless of index
                mstore(124, ERC20_TRANSFER_SIG)
                mstore(128, target)// destination
                mstore(160, sub(add(_amountOut0,_amountOut1),revenue))// amount
                pop(call(gas(),WETH_ADDRESS,0,124,68,0,0))

                //do frist swap
                let amountOutIndexes := shr(248,calldataload(167)) //1 byte

                let mask := 1
                let amountOutIndex := iszero(and(mask,amountOutIndexes))

                //swap call
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
                pop(call(gas(), target, 0, 0x7c, 0xa4, 0, 0))

                //increment position by i*52. 52 is the size of the data for each swap
                currPos := add(currPos, swapDataSize)

                
                let swapCountMinus1 := shr(248,calldataload(166)) //197 1 byte we subtract 1 off chain to save gas as the last swap has custom logic

                for { let i := 1 } lt(i, swapCountMinus1) { i := add(i, 1) } 
                {  
                    //mask checks if bit i is 0 or 1
                    mask := shl(i,1)
                    amountOutIndex := iszero(and(mask,amountOutIndexes))

                    target := shr(96, calldataload(currPos))//20 bytes
                    amountOut := shr(intRightShift,calldataload(add(currPos,20)))
                    to := shr(96, calldataload(add(currPos,swapDataSize)))

                    //swap call
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
                    pop(call(gas(), target, 0, 0x7c, 0xa4, 0, 0))

                    //increment position by swapDataSize
                    currPos := add(currPos, swapDataSize)

                }

                mask := shl(swapCountMinus1,1)
                amountOutIndex := iszero(and(mask,amountOutIndexes))

                target := shr(96, calldataload(currPos))//20 bytes
                amountOut := shr(intRightShift,calldataload(add(currPos,20)))
                //let to := shr(96, calldataload(add(currPos,52)))

                //swap call
                mstore(0x7c, PAIR_SWAP_SIG)
                switch amountOutIndex // tokenOutNo == 0 ? ....
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
                pop(call(gas(), target, 0, 0x7c, 0xa4, 0, 0))
            }
        }
    }