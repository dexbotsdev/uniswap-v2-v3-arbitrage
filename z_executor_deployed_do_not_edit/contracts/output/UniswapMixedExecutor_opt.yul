/// @use-src 0:"sources/UniswapMixedExecutor.sol"
object "UniswapMixedExecutor_1550" {
    code {
        {
            /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
            let _1 := memoryguard(0xa0)
            mstore(64, _1)
            /// @src 0:2338:2389  "payable(0x298806238A1b5DAF30Bf290D01b8F92036eD36F6)"
            mstore(128, /** @src 0:2346:2388  "0x298806238A1b5DAF30Bf290D01b8F92036eD36F6" */ 0x298806238a1b5daf30bf290d01b8f92036ed36f6)
            /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
            let _2 := datasize("UniswapMixedExecutor_1550_deployed")
            codecopy(_1, dataoffset("UniswapMixedExecutor_1550_deployed"), _2)
            setimmutable(_1, "184", mload(/** @src 0:2338:2389  "payable(0x298806238A1b5DAF30Bf290D01b8F92036eD36F6)" */ 128))
            /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
            return(_1, _2)
        }
    }
    /// @use-src 0:"sources/UniswapMixedExecutor.sol"
    object "UniswapMixedExecutor_1550_deployed" {
        code {
            {
                /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                mstore(64, memoryguard(0x80))
                fun()
                stop()
            }
            /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
            function finalize_allocation_6340(memPtr)
            {
                if gt(memPtr, 0xffffffffffffffff)
                {
                    mstore(0, 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                    mstore(4, 0x41)
                    revert(0, 0x24)
                }
                mstore(64, memPtr)
            }
            function finalize_allocation_8896(memPtr)
            {
                let newFreePtr := add(memPtr, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ 128)
                /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                if or(gt(newFreePtr, 0xffffffffffffffff), lt(newFreePtr, memPtr))
                {
                    mstore(0, 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                    mstore(4, 0x41)
                    revert(0, 0x24)
                }
                mstore(64, newFreePtr)
            }
            function finalize_allocation_8900(memPtr)
            {
                let newFreePtr := add(memPtr, /** @src 0:20342:20383  "WETH.transfer(firstTarget, firstAmountIn)" */ 32)
                /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                if or(gt(newFreePtr, 0xffffffffffffffff), lt(newFreePtr, memPtr))
                {
                    mstore(0, 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                    mstore(4, 0x41)
                    revert(0, 0x24)
                }
                mstore(64, newFreePtr)
            }
            function finalize_allocation(memPtr, size)
            {
                let newFreePtr := add(memPtr, and(add(size, 31), 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0))
                if or(gt(newFreePtr, 0xffffffffffffffff), lt(newFreePtr, memPtr))
                {
                    mstore(0, 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                    mstore(4, 0x41)
                    revert(0, 0x24)
                }
                mstore(64, newFreePtr)
            }
            function checked_add_uint256(x, y) -> sum
            {
                sum := add(x, y)
                if gt(x, sum)
                {
                    mstore(0, 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                    mstore(4, 0x11)
                    revert(0, 0x24)
                }
            }
            function convert_bytes_to_fixedbytes_from_bytes_calldata_to_bytes20(array, len) -> value
            {
                let _1 := calldataload(array)
                let _2 := 0xffffffffffffffffffffffffffffffffffffffff000000000000000000000000
                value := and(_1, _2)
                if lt(len, 20)
                {
                    value := and(and(_1, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ shl(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ shl(3, sub(20, len)), _2)), _2)
                }
            }
            function checked_mul_uint8(x) -> product
            {
                let product_raw := mul(and(x, 0xff), 20)
                product := and(product_raw, 0xff)
                if iszero(eq(product, product_raw))
                {
                    mstore(0, 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                    mstore(4, 0x11)
                    revert(0, 0x24)
                }
            }
            /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
            function require_helper_stringliteral_175a(condition)
            {
                if iszero(condition)
                {
                    let memPtr := /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ mload(64)
                    /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                    mstore(memPtr, 0x08c379a000000000000000000000000000000000000000000000000000000000)
                    mstore(add(memPtr, 4), 32)
                    mstore(add(memPtr, 36), 14)
                    mstore(add(memPtr, 68), "invalid sender")
                    revert(memPtr, 100)
                }
            }
            function checked_mul_uint256(x, y) -> product
            {
                product := mul(x, y)
                if iszero(or(iszero(x), eq(y, div(product, x))))
                {
                    /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                    mstore(0, 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                    mstore(4, 0x11)
                    revert(0, 0x24)
                }
            }
            /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
            function convert_bytes_to_fixedbytes_from_bytes_calldata_to_bytes32(array, len) -> value
            {
                value := calldataload(array)
                if lt(len, 32)
                {
                    value := and(value, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ shl(/** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ shl(3, sub(32, len)), 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff))
                }
            }
            function copy_memory_to_memory_with_cleanup(src, dst, length)
            {
                let i := 0
                for { } lt(i, length) { i := add(i, 32) }
                {
                    mstore(add(dst, i), mload(add(src, i)))
                }
                mstore(add(dst, length), 0)
            }
            function abi_encode_bytes(value, pos) -> end
            {
                let length := mload(value)
                mstore(pos, length)
                copy_memory_to_memory_with_cleanup(add(value, 0x20), add(pos, 0x20), length)
                end := add(add(pos, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ and(add(/** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ length, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 31), 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0)), /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 0x20)
            }
            function checked_add_uint8(x) -> sum
            {
                sum := add(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ and(/** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ x, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0xff), 1)
                /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                if gt(sum, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0xff)
                /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                {
                    /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                    mstore(0, 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                    mstore(4, 0x11)
                    revert(0, 0x24)
                }
            }
            /// @src 0:3350:3360  "4295128740"
            function abi_decode_available_length_bytes(src, length, end) -> array
            {
                /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                if gt(length, 0xffffffffffffffff)
                {
                    /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                    mstore(/** @src -1:-1:-1 */ 0, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                    mstore(4, 0x41)
                    revert(/** @src -1:-1:-1 */ 0, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                }
                /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                let memPtr := /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ mload(64)
                /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                finalize_allocation(memPtr, add(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ and(add(/** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ length, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 31), 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0), /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 0x20))
                /// @src 0:3350:3360  "4295128740"
                array := memPtr
                mstore(memPtr, length)
                if gt(add(src, length), end)
                {
                    /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                    revert(/** @src -1:-1:-1 */ 0, 0)
                }
                /// @src 0:3350:3360  "4295128740"
                calldatacopy(add(memPtr, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 0x20), /** @src 0:3350:3360  "4295128740" */ src, length)
                mstore(add(add(memPtr, length), /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 0x20), /** @src -1:-1:-1 */ 0)
            }
            /// @ast-id 1372 @src 0:3735:22975  "fallback() external payable{..."
            function fun()
            {
                /// @src 0:4071:4090  "msg.data.length ==0"
                let _1 := /** @src 0:4071:4079  "msg.data" */ 0
                /// @src 0:4067:4156  "if (msg.data.length ==0){..."
                if /** @src 0:4071:4090  "msg.data.length ==0" */ iszero(/** @src 0:4071:4079  "msg.data" */ calldatasize())
                /// @src 0:4067:4156  "if (msg.data.length ==0){..."
                {
                    /// @src 0:4138:4145  "return;"
                    leave
                }
                /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                if gt(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:4071:4079  "msg.data" */ calldatasize())
                /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                {
                    revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                }
                /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                let value := and(calldataload(/** @src 0:4071:4079  "msg.data" */ _1), /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ 0xffffffff00000000000000000000000000000000000000000000000000000000)
                /// @src 0:4331:4355  "uint mainPayloadStartPos"
                let var_mainPayloadStartPos := /** @src 0:4071:4079  "msg.data" */ _1
                /// @src 0:4331:4355  "uint mainPayloadStartPos"
                var_mainPayloadStartPos := /** @src 0:4071:4079  "msg.data" */ _1
                /// @src 0:4366:4390  "bool v3SwapsDone = false"
                let var_v3SwapsDone := /** @src 0:4071:4079  "msg.data" */ _1
                /// @src 0:4492:11566  "if (selector == Uniswap_V2_CALLBACK_SIG){            ..."
                switch /** @src 0:4496:4531  "selector == Uniswap_V2_CALLBACK_SIG" */ eq(value, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ 0x10d1e85c00000000000000000000000000000000000000000000000000000000)
                case /** @src 0:4492:11566  "if (selector == Uniswap_V2_CALLBACK_SIG){            ..." */ 0 {
                    /// @src 0:6853:11566  "if (selector == Uniswap_V3_CALLBACK_SIG){//***V3 CALLBACK***..."
                    switch /** @src 0:6857:6892  "selector == Uniswap_V3_CALLBACK_SIG" */ eq(value, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ 0xfa461e3300000000000000000000000000000000000000000000000000000000)
                    case /** @src 0:6853:11566  "if (selector == Uniswap_V3_CALLBACK_SIG){//***V3 CALLBACK***..." */ 0 {
                        /// @src 0:10898:11566  "if (selector == WITHDRAW_SIG) {..."
                        switch /** @src 0:10902:10926  "selector == WITHDRAW_SIG" */ eq(value, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ 0x3ccfd60b00000000000000000000000000000000000000000000000000000000)
                        case /** @src 0:10898:11566  "if (selector == WITHDRAW_SIG) {..." */ 0 {
                            /// @src 0:11446:11469  "mainPayloadStartPos = 0"
                            var_mainPayloadStartPos := /** @src 0:4071:4079  "msg.data" */ _1
                        }
                        default /// @src 0:10898:11566  "if (selector == WITHDRAW_SIG) {..."
                        {
                            /// @src 0:11120:11149  "WETH.balanceOf(address(this))"
                            let _2 := /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ mload(64)
                            /// @src 0:11120:11149  "WETH.balanceOf(address(this))"
                            mstore(_2, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ 0x70a0823100000000000000000000000000000000000000000000000000000000)
                            /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                            mstore(/** @src 0:11120:11149  "WETH.balanceOf(address(this))" */ add(_2, /** @src 0:4235:4236  "4" */ 0x04), /** @src 0:11143:11147  "this" */ address())
                            /// @src 0:2484:2526  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                            let _3 := 0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2
                            /// @src 0:11120:11149  "WETH.balanceOf(address(this))"
                            let _4 := staticcall(gas(), /** @src 0:2484:2526  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ _3, /** @src 0:11120:11149  "WETH.balanceOf(address(this))" */ _2, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 36, /** @src 0:11120:11149  "WETH.balanceOf(address(this))" */ _2, 32)
                            if iszero(_4)
                            {
                                /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                                let pos := /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ mload(64)
                                /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                                returndatacopy(pos, /** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ returndatasize())
                                revert(pos, returndatasize())
                            }
                            /// @src 0:11120:11149  "WETH.balanceOf(address(this))"
                            let expr := /** @src 0:4071:4079  "msg.data" */ _1
                            /// @src 0:11120:11149  "WETH.balanceOf(address(this))"
                            if _4
                            {
                                let _5 := 32
                                if gt(_5, returndatasize()) { _5 := returndatasize() }
                                finalize_allocation(_2, _5)
                                /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                                if slt(sub(/** @src 0:11120:11149  "WETH.balanceOf(address(this))" */ add(_2, _5), /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ _2), /** @src 0:11120:11149  "WETH.balanceOf(address(this))" */ 32)
                                /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                                {
                                    /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                                    revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                                }
                                /// @src 0:11120:11149  "WETH.balanceOf(address(this))"
                                expr := /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ mload(_2)
                            }
                            /// @src 0:11164:11245  "if (wethBalance > 0){..."
                            if /** @src 0:11168:11183  "wethBalance > 0" */ iszero(iszero(expr))
                            /// @src 0:11164:11245  "if (wethBalance > 0){..."
                            {
                                /// @src 0:11203:11229  "WETH.withdraw(wethBalance)"
                                if iszero(extcodesize(/** @src 0:2484:2526  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ _3))
                                /// @src 0:11203:11229  "WETH.withdraw(wethBalance)"
                                {
                                    /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                                    revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                                }
                                /// @src 0:11203:11229  "WETH.withdraw(wethBalance)"
                                let _6 := /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ mload(64)
                                /// @src 0:11203:11229  "WETH.withdraw(wethBalance)"
                                mstore(_6, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ 0x2e1a7d4d00000000000000000000000000000000000000000000000000000000)
                                /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                                mstore(/** @src 0:11203:11229  "WETH.withdraw(wethBalance)" */ add(_6, /** @src 0:4235:4236  "4" */ 0x04), /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ expr)
                                /// @src 0:11203:11229  "WETH.withdraw(wethBalance)"
                                let _7 := call(gas(), /** @src 0:2484:2526  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ _3, /** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:11203:11229  "WETH.withdraw(wethBalance)" */ _6, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 36, /** @src 0:11203:11229  "WETH.withdraw(wethBalance)" */ _6, /** @src 0:4071:4079  "msg.data" */ _1)
                                /// @src 0:11203:11229  "WETH.withdraw(wethBalance)"
                                if iszero(_7)
                                {
                                    /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                                    let pos_1 := /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ mload(64)
                                    /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                                    returndatacopy(pos_1, /** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ returndatasize())
                                    revert(pos_1, returndatasize())
                                }
                                /// @src 0:11203:11229  "WETH.withdraw(wethBalance)"
                                if _7
                                {
                                    finalize_allocation_6340(_6)
                                    /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                                    _1 := /** @src 0:4071:4079  "msg.data" */ _1
                                }
                            }
                            /// @src 0:11274:11295  "address(this).balance"
                            let expr_1 := selfbalance()
                            /// @src 0:11259:11296  "owner.transfer(address(this).balance)"
                            let _8 := _1
                            if iszero(expr_1) { _8 := 2300 }
                            if iszero(call(_8, /** @src 0:3159:3169  "0x3ccfd60b" */ and(/** @src 0:11259:11264  "owner" */ loadimmutable("184"), /** @src 0:3159:3169  "0x3ccfd60b" */ 0xffffffffffffffffffffffffffffffffffffffff), /** @src 0:11259:11296  "owner.transfer(address(this).balance)" */ expr_1, /** @src 0:4071:4079  "msg.data" */ _1, _1, _1, _1))
                            /// @src 0:11259:11296  "owner.transfer(address(this).balance)"
                            {
                                /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                                let pos_2 := /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ mload(64)
                                /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                                returndatacopy(pos_2, /** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ returndatasize())
                                revert(pos_2, returndatasize())
                            }
                            /// @src 0:11311:11318  "return;"
                            leave
                        }
                    }
                    default /// @src 0:6853:11566  "if (selector == Uniswap_V3_CALLBACK_SIG){//***V3 CALLBACK***..."
                    {
                        /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                        if iszero(lt(/** @src 0:7475:7478  "133" */ 0x85, /** @src 0:4071:4079  "msg.data" */ calldatasize()))
                        /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                        {
                            mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                            mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x32)
                            revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                        }
                        /// @src 0:7466:7479  "msg.data[133]"
                        let _9 := calldataload(/** @src 0:7475:7478  "133" */ 0x85)
                        /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                        let _10 := shr(248, _9)
                        /// @src 0:7732:7746  "targetByte & 7"
                        let expr_2 := and(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ _10, /** @src 0:7745:7746  "7" */ 0x07)
                        /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                        if iszero(lt(/** @src 0:7876:7879  "134" */ 0x86, /** @src 0:4071:4079  "msg.data" */ calldatasize()))
                        /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                        {
                            mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                            mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x32)
                            revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                        }
                        /// @src 0:7867:7880  "msg.data[134]"
                        let _11 := calldataload(/** @src 0:7876:7879  "134" */ 0x86)
                        /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                        let _12 := shr(251, _11)
                        /// @src 0:7966:7980  "targetByte & 7"
                        let expr_3 := and(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ shr(248, _11), /** @src 0:7745:7746  "7" */ 0x07)
                        /// @src 0:7997:8074  "if (callbackV3SwapsDone){..."
                        if /** @src 0:7597:7618  "targetByte & 32 == 32" */ eq(/** @src 0:7597:7612  "targetByte & 32" */ and(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ _10, /** @src 0:7610:7612  "32" */ 0x20), 0x20)
                        /// @src 0:7997:8074  "if (callbackV3SwapsDone){..."
                        {
                            /// @src 0:8040:8058  "v3SwapsDone = true"
                            var_v3SwapsDone := /** @src 0:8054:8058  "true" */ 0x01
                        }
                        /// @src 0:8115:8131  "getFee(feeIndex)"
                        let expr_4 := fun_getFee(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ and(shr(251, _9), /** @src 0:7681:7682  "3" */ 0x03))
                        /// @src 0:8193:8211  "uint currPos = 135"
                        let var_currPos := /** @src 0:8208:8211  "135" */ 0x87
                        /// @src 0:8288:8303  "address factory"
                        let var_factory := _1
                        /// @src 0:8318:8625  "if (factoryIndex == 1){ //factory index = 1 means custom address..."
                        switch /** @src 0:8322:8339  "factoryIndex == 1" */ eq(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ expr_2, /** @src 0:8338:8339  "1" */ 0x01)
                        case /** @src 0:8318:8625  "if (factoryIndex == 1){ //factory index = 1 means custom address..." */ 0 {
                            /// @src 0:8573:8609  "factory = getV3Factory(factoryIndex)"
                            var_factory := /** @src 0:8583:8609  "getV3Factory(factoryIndex)" */ fun_getV3Factory(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ expr_2)
                        }
                        default /// @src 0:8318:8625  "if (factoryIndex == 1){ //factory index = 1 means custom address..."
                        {
                            /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                            if gt(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 155, /** @src 0:4071:4079  "msg.data" */ calldatasize())
                            /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                            {
                                revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                            }
                            /// @src 0:8400:8456  "factory = address(bytes20(msg.data[currPos:currPos+20]))"
                            var_factory := /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ shr(96, calldataload(/** @src 0:8208:8211  "135" */ var_currPos))
                            /// @src 0:8475:8488  "currPos += 20"
                            var_currPos := /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 155
                        }
                        if iszero(lt(var_currPos, /** @src 0:4071:4079  "msg.data" */ calldatasize()))
                        /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                        {
                            mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                            mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x32)
                            revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                        }
                        let sum := add(var_currPos, /** @src 0:9019:9020  "2" */ 0x02)
                        /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                        if gt(var_currPos, sum)
                        {
                            mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                            mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x11)
                            revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                        }
                        if iszero(lt(sum, /** @src 0:4071:4079  "msg.data" */ calldatasize()))
                        /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                        {
                            mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                            mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x32)
                            revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                        }
                        /// @src 0:9058:9088  "uint8(msg.data[currPos+2]) & 7"
                        let expr_5 := and(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ shr(248, /** @src 0:9064:9083  "msg.data[currPos+2]" */ calldataload(sum)), /** @src 0:7745:7746  "7" */ 0x07)
                        /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                        let sum_1 := add(var_currPos, /** @src 0:4235:4236  "4" */ 0x04)
                        /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                        if gt(var_currPos, sum_1)
                        {
                            mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                            mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x11)
                            revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                        }
                        /// @src 0:9244:9269  "currPos + 4 + intByteSize"
                        let expr_6 := checked_add_uint256(sum_1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ shr(248, /** @src 0:8701:8718  "msg.data[currPos]" */ calldataload(var_currPos)))
                        /// @src 0:9415:9429  "address token0"
                        let var_token0 := _1
                        /// @src 0:9444:9873  "if (token0Index == tokensLength){//if token0Index == tokensLength, set to weth..."
                        switch /** @src 0:9448:9475  "token0Index == tokensLength" */ eq(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ _12, /** @src 0:9448:9475  "token0Index == tokensLength" */ expr_5)
                        case /** @src 0:9444:9873  "if (token0Index == tokensLength){//if token0Index == tokensLength, set to weth..." */ 0 {
                            /// @src 0:9791:9819  "tokensPos + token0Index * 20"
                            let expr_7 := checked_add_uint256(expr_6, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ and(/** @src 0:9803:9819  "token0Index * 20" */ checked_mul_uint8(_12), /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0xff))
                            /// @src 0:9821:9849  "tokensPos + token0Index * 20"
                            let expr_8 := checked_add_uint256(expr_6, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ and(/** @src 0:9833:9849  "token0Index * 20" */ checked_mul_uint8(_12), /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0xff))
                            let sum_2 := add(expr_8, /** @src 0:9817:9819  "20" */ 0x14)
                            /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                            if gt(expr_8, sum_2)
                            {
                                mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                                mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x11)
                                revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                            }
                            /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                            if gt(expr_7, sum_2)
                            {
                                revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                            }
                            /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                            if gt(sum_2, /** @src 0:4071:4079  "msg.data" */ calldatasize())
                            /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                            {
                                revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                            }
                            /// @src 0:9757:9857  "token0 = address(bytes20(msg.data[tokensPos + token0Index * 20: tokensPos + token0Index * 20 + 20]))"
                            var_token0 := /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ shr(96, /** @src 0:9774:9856  "bytes20(msg.data[tokensPos + token0Index * 20: tokensPos + token0Index * 20 + 20])" */ convert_bytes_to_fixedbytes_from_bytes_calldata_to_bytes20(expr_7, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ add(sub(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ expr_8, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ expr_7), /** @src 0:9817:9819  "20" */ 0x14)))
                        }
                        default /// @src 0:9444:9873  "if (token0Index == tokensLength){//if token0Index == tokensLength, set to weth..."
                        {
                            /// @src 0:9695:9716  "token0 = WETH_ADDRESS"
                            var_token0 := /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2
                        }
                        /// @src 0:9961:9975  "address token1"
                        let var_token1 := _1
                        /// @src 0:9990:10374  "if (token1Index == tokensLength){..."
                        switch /** @src 0:9994:10021  "token1Index == tokensLength" */ eq(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ expr_3, /** @src 0:9994:10021  "token1Index == tokensLength" */ expr_5)
                        case /** @src 0:9990:10374  "if (token1Index == tokensLength){..." */ 0 {
                            /// @src 0:10292:10320  "tokensPos + token1Index * 20"
                            let expr_9 := checked_add_uint256(expr_6, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ and(/** @src 0:10304:10320  "token1Index * 20" */ checked_mul_uint8(expr_3), /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0xff))
                            /// @src 0:10322:10350  "tokensPos + token1Index * 20"
                            let expr_10 := checked_add_uint256(expr_6, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ and(/** @src 0:10334:10350  "token1Index * 20" */ checked_mul_uint8(expr_3), /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0xff))
                            let sum_3 := add(expr_10, /** @src 0:10318:10320  "20" */ 0x14)
                            /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                            if gt(expr_10, sum_3)
                            {
                                mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                                mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x11)
                                revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                            }
                            /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                            if gt(expr_9, sum_3)
                            {
                                revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                            }
                            /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                            if gt(sum_3, /** @src 0:4071:4079  "msg.data" */ calldatasize())
                            /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                            {
                                revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                            }
                            /// @src 0:10258:10358  "token1 = address(bytes20(msg.data[tokensPos + token1Index * 20: tokensPos + token1Index * 20 + 20]))"
                            var_token1 := /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ shr(96, /** @src 0:10275:10357  "bytes20(msg.data[tokensPos + token1Index * 20: tokensPos + token1Index * 20 + 20])" */ convert_bytes_to_fixedbytes_from_bytes_calldata_to_bytes20(expr_9, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ add(sub(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ expr_10, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ expr_9), /** @src 0:10318:10320  "20" */ 0x14)))
                        }
                        default /// @src 0:9990:10374  "if (token1Index == tokensLength){..."
                        {
                            /// @src 0:10196:10217  "token1 = WETH_ADDRESS"
                            var_token1 := /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2
                        }
                        /// @src 0:23640:23671  "abi.encode(token0, token1, fee)"
                        let expr_mpos := /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ mload(64)
                        /// @src 0:23640:23671  "abi.encode(token0, token1, fee)"
                        let _13 := add(expr_mpos, /** @src 0:7610:7612  "32" */ 0x20)
                        /// @src 0:3159:3169  "0x3ccfd60b"
                        let _14 := 0xffffffffffffffffffffffffffffffffffffffff
                        /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                        mstore(_13, /** @src 0:3159:3169  "0x3ccfd60b" */ and(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ var_token0, /** @src 0:3159:3169  "0x3ccfd60b" */ _14))
                        /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                        mstore(/** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ add(/** @src 0:23640:23671  "abi.encode(token0, token1, fee)" */ expr_mpos, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ 64), /** @src 0:3159:3169  "0x3ccfd60b" */ and(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ var_token1, /** @src 0:3159:3169  "0x3ccfd60b" */ _14))
                        /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                        mstore(add(/** @src 0:23640:23671  "abi.encode(token0, token1, fee)" */ expr_mpos, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ 96), and(expr_4, 0xffffff))
                        /// @src 0:23640:23671  "abi.encode(token0, token1, fee)"
                        mstore(expr_mpos, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ 96)
                        /// @src 0:23640:23671  "abi.encode(token0, token1, fee)"
                        finalize_allocation_8896(expr_mpos)
                        /// @src 0:23630:23672  "keccak256(abi.encode(token0, token1, fee))"
                        let expr_11 := keccak256(/** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ _13, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ mload(/** @src 0:23630:23672  "keccak256(abi.encode(token0, token1, fee))" */ expr_mpos))
                        /// @src 0:23531:23783  "abi.encodePacked(..."
                        let expr_mpos_1 := /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ mload(64)
                        /// @src 0:23531:23783  "abi.encodePacked(..."
                        let _15 := add(expr_mpos_1, /** @src 0:7610:7612  "32" */ 0x20)
                        /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                        mstore(_15, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0xff00000000000000000000000000000000000000000000000000000000000000)
                        /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                        mstore(add(/** @src 0:23531:23783  "abi.encodePacked(..." */ expr_mpos_1, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ 33), and(shl(96, var_factory), 0xffffffffffffffffffffffffffffffffffffffff000000000000000000000000))
                        /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                        mstore(/** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ add(/** @src 0:23531:23783  "abi.encodePacked(..." */ expr_mpos_1, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ 53), /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ expr_11)
                        /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                        mstore(add(/** @src 0:23531:23783  "abi.encodePacked(..." */ expr_mpos_1, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ 85), 0xe34f199b19b2b4f47f68442619d555527d244f78a3297ea89325f843f87b8b54)
                        /// @src 0:23531:23783  "abi.encodePacked(..."
                        mstore(expr_mpos_1, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ 85)
                        /// @src 0:23531:23783  "abi.encodePacked(..."
                        finalize_allocation_8896(expr_mpos_1)
                        /// @src 0:10580:10678  "require(uniswapV3CallbackVerification(msg.sender, factory, token0, token1, fee), \"invalid sender\")"
                        require_helper_stringliteral_175a(/** @src 0:23882:23906  "pool == poolToBeVerified" */ eq(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ and(/** @src 0:23503:23798  "keccak256(..." */ keccak256(/** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ _15, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ mload(/** @src 0:23503:23798  "keccak256(..." */ expr_mpos_1)), /** @src 0:3159:3169  "0x3ccfd60b" */ _14), /** @src 0:10618:10628  "msg.sender" */ caller()))
                        /// @src 0:10851:10880  "mainPayloadStartPos = currPos"
                        var_mainPayloadStartPos := var_currPos
                    }
                }
                default /// @src 0:4492:11566  "if (selector == Uniswap_V2_CALLBACK_SIG){            ..."
                {
                    /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                    if iszero(lt(/** @src 0:5136:5139  "164" */ 0xa4, /** @src 0:4071:4079  "msg.data" */ calldatasize()))
                    /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                    {
                        mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                        mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x32)
                        revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                    }
                    /// @src 0:5127:5140  "msg.data[164]"
                    let _16 := calldataload(/** @src 0:5136:5139  "164" */ 0xa4)
                    /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                    let _17 := shr(248, _16)
                    let _18 := and(shr(251, _16), 7)
                    /// @src 0:5303:5317  "targetByte & 7"
                    let _19 := and(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ _17, 7)
                    /// @src 0:5396:5414  "uint currPos = 165"
                    let var_currPos_1 := /** @src 0:5411:5414  "165" */ 0xa5
                    /// @src 0:5431:5510  "if (v2CallbackV3SwapsDone){..."
                    if /** @src 0:5185:5206  "targetByte & 64 == 64" */ eq(/** @src 0:5185:5200  "targetByte & 64" */ and(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ _17, /** @src 0:5198:5200  "64" */ 0x40), 0x40)
                    /// @src 0:5431:5510  "if (v2CallbackV3SwapsDone){..."
                    {
                        /// @src 0:5476:5494  "v3SwapsDone = true"
                        var_v3SwapsDone := /** @src 0:5490:5494  "true" */ 0x01
                    }
                    /// @src 0:5526:5541  "address factory"
                    let var_factory_1 := _1
                    /// @src 0:5556:5840  "if (v2FactoryIndex == 1){ //1 means custom address..."
                    switch /** @src 0:5560:5579  "v2FactoryIndex == 1" */ eq(_18, /** @src 0:5578:5579  "1" */ 0x01)
                    case /** @src 0:5556:5840  "if (v2FactoryIndex == 1){ //1 means custom address..." */ 0 {
                        /// @src 0:5786:5824  "factory = getV2Factory(v2FactoryIndex)"
                        var_factory_1 := /** @src 0:5796:5824  "getV2Factory(v2FactoryIndex)" */ fun_getV2Factory(_18)
                    }
                    default /// @src 0:5556:5840  "if (v2FactoryIndex == 1){ //1 means custom address..."
                    {
                        /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                        if gt(/** @src 0:5663:5666  "185" */ 0xb9, /** @src 0:4071:4079  "msg.data" */ calldatasize())
                        /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                        {
                            revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                        }
                        /// @src 0:5624:5669  "factory = address(bytes20(msg.data[165:185]))"
                        var_factory_1 := /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ shr(96, calldataload(/** @src 0:5411:5414  "165" */ var_currPos_1))
                        /// @src 0:5688:5701  "currPos += 20"
                        var_currPos_1 := /** @src 0:5663:5666  "185" */ 0xb9
                    }
                    /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                    if iszero(lt(var_currPos_1, /** @src 0:4071:4079  "msg.data" */ calldatasize()))
                    /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                    {
                        mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                        mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x32)
                        revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                    }
                    let sum_4 := add(var_currPos_1, /** @src 0:4235:4236  "4" */ 0x04)
                    /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                    if gt(var_currPos_1, sum_4)
                    {
                        mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                        mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x11)
                        revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                    }
                    /// @src 0:6243:6268  "currPos + 4 + intByteSize"
                    let expr_12 := checked_add_uint256(sum_4, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ shr(248, /** @src 0:5965:5982  "msg.data[currPos]" */ calldataload(var_currPos_1)))
                    /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                    let product := mul(_19, /** @src 0:6361:6363  "20" */ 0x14)
                    /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                    if iszero(or(iszero(_19), eq(/** @src 0:6361:6363  "20" */ 0x14, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ div(product, _19))))
                    {
                        /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                        mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                        mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x11)
                        revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                    }
                    /// @src 0:6329:6363  "tokensPos + v2OtherTokenIndex * 20"
                    let expr_13 := checked_add_uint256(expr_12, /** @src 0:6341:6363  "v2OtherTokenIndex * 20" */ product)
                    /// @src 0:6365:6399  "tokensPos + v2OtherTokenIndex * 20"
                    let _20 := checked_add_uint256(expr_12, /** @src 0:6377:6399  "v2OtherTokenIndex * 20" */ product)
                    /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                    let sum_5 := add(_20, /** @src 0:6361:6363  "20" */ 0x14)
                    /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                    if gt(_20, sum_5)
                    {
                        mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                        mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x11)
                        revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                    }
                    /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                    if gt(expr_13, sum_5)
                    {
                        revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                    }
                    /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                    if gt(sum_5, /** @src 0:4071:4079  "msg.data" */ calldatasize())
                    /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                    {
                        revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                    }
                    /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                    let _21 := shr(96, /** @src 0:6312:6406  "bytes20(msg.data[tokensPos + v2OtherTokenIndex * 20: tokensPos + v2OtherTokenIndex * 20 + 20])" */ convert_bytes_to_fixedbytes_from_bytes_calldata_to_bytes20(expr_13, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ add(sub(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ _20, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ expr_13), /** @src 0:6361:6363  "20" */ 0x14)))
                    /// @src 0:6465:6494  "address token0 = WETH_ADDRESS"
                    let var_token0_1 := /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2
                    /// @src 0:6509:6536  "address token1 = otherToken"
                    let var_token1_1 := _21
                    /// @src 0:6551:6607  "if (token0 > token1) (token0, token1) = (token1, token0)"
                    if /** @src 0:6555:6570  "token0 > token1" */ gt(/** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ var_token0_1, /** @src 0:3159:3169  "0x3ccfd60b" */ _21)
                    /// @src 0:6551:6607  "if (token0 > token1) (token0, token1) = (token1, token0)"
                    {
                        /// @src 0:6572:6607  "(token0, token1) = (token1, token0)"
                        var_token1_1 := /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ var_token0_1
                        /// @src 0:6572:6607  "(token0, token1) = (token1, token0)"
                        var_token0_1 := _21
                    }
                    /// @src 0:24483:24515  "abi.encodePacked(token0, token1)"
                    let expr_mpos_2 := /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ mload(/** @src 0:5198:5200  "64" */ 0x40)
                    /// @src 0:24483:24515  "abi.encodePacked(token0, token1)"
                    let _22 := add(expr_mpos_2, 0x20)
                    /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                    let _23 := 0xffffffffffffffffffffffffffffffffffffffff000000000000000000000000
                    mstore(_22, and(shl(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 96, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ var_token0_1), _23))
                    mstore(add(/** @src 0:24483:24515  "abi.encodePacked(token0, token1)" */ expr_mpos_2, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ 52), and(shl(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 96, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ var_token1_1), _23))
                    /// @src 0:24483:24515  "abi.encodePacked(token0, token1)"
                    mstore(expr_mpos_2, 40)
                    /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                    let newFreePtr := add(expr_mpos_2, 96)
                    if or(gt(newFreePtr, 0xffffffffffffffff), lt(newFreePtr, expr_mpos_2))
                    {
                        mstore(/** @src 0:4071:4079  "msg.data" */ 0, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                        mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x41)
                        revert(/** @src 0:4071:4079  "msg.data" */ 0, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                    }
                    mstore(/** @src 0:5198:5200  "64" */ 0x40, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ newFreePtr)
                    /// @src 0:24473:24516  "keccak256(abi.encodePacked(token0, token1))"
                    let expr_14 := keccak256(/** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ _22, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ mload(/** @src 0:24473:24516  "keccak256(abi.encodePacked(token0, token1))" */ expr_mpos_2))
                    /// @src 0:24398:24611  "abi.encodePacked(..."
                    let _24 := add(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ expr_mpos_2, /** @src 0:24398:24611  "abi.encodePacked(..." */ 128)
                    /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                    mstore(_24, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0xff00000000000000000000000000000000000000000000000000000000000000)
                    /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                    mstore(add(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ expr_mpos_2, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ 129), and(shl(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 96, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ var_factory_1), _23))
                    /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                    mstore(/** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ add(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ expr_mpos_2, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ 149), /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ expr_14)
                    /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                    mstore(add(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ expr_mpos_2, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ 181), 0x96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f)
                    /// @src 0:24398:24611  "abi.encodePacked(..."
                    mstore(newFreePtr, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ 85)
                    /// @src 0:24398:24611  "abi.encodePacked(..."
                    finalize_allocation_8896(newFreePtr)
                    /// @src 0:6660:6751  "require(uniswapV2CallbackVerification(msg.sender, factory,token0,token1), \"invalid sender\")"
                    require_helper_stringliteral_175a(/** @src 0:24688:24712  "pair == poolToBeVerified" */ eq(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ and(/** @src 0:24388:24612  "keccak256(abi.encodePacked(..." */ keccak256(/** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ _24, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ mload(/** @src 0:24388:24612  "keccak256(abi.encodePacked(..." */ newFreePtr)), /** @src 0:3159:3169  "0x3ccfd60b" */ 0xffffffffffffffffffffffffffffffffffffffff), /** @src 0:6698:6708  "msg.sender" */ caller()))
                    /// @src 0:6776:6805  "mainPayloadStartPos = currPos"
                    var_mainPayloadStartPos := var_currPos_1
                }
                /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                if iszero(lt(var_mainPayloadStartPos, /** @src 0:4071:4079  "msg.data" */ calldatasize()))
                /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                {
                    mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                    mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x32)
                    revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                }
                /// @src 0:11701:11730  "msg.data[mainPayloadStartPos]"
                let _25 := calldataload(var_mainPayloadStartPos)
                /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                let _26 := and(shr(245, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ _25), /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 2040)
                if iszero(or(iszero(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ shr(248, _25)), /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ eq(/** @src 0:11781:11782  "8" */ 0x08, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ div(_26, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ shr(248, _25)))))
                /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                {
                    /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                    mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                    mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x11)
                    revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                }
                /// @src 0:11763:11766  "256"
                let _27 := 0x0100
                /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                if gt(sub(/** @src 0:11763:11766  "256" */ _27, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ _26), /** @src 0:11763:11766  "256" */ _27)
                /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                {
                    /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                    mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                    mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x11)
                    revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                }
                let sum_6 := add(var_mainPayloadStartPos, /** @src 0:12038:12039  "1" */ 0x01)
                /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                if gt(var_mainPayloadStartPos, sum_6)
                {
                    mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                    mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x11)
                    revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                }
                if iszero(lt(sum_6, /** @src 0:4071:4079  "msg.data" */ calldatasize()))
                /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                {
                    mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                    mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x32)
                    revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                }
                /// @src 0:12009:12040  "msg.data[mainPayloadStartPos+1]"
                let _28 := calldataload(sum_6)
                /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                let sum_7 := add(var_mainPayloadStartPos, /** @src 0:12511:12512  "2" */ 0x02)
                /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                if gt(var_mainPayloadStartPos, sum_7)
                {
                    mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                    mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x11)
                    revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                }
                if iszero(lt(sum_7, /** @src 0:4071:4079  "msg.data" */ calldatasize()))
                /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                {
                    mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                    mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x32)
                    revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                }
                /// @src 0:12482:12513  "msg.data[mainPayloadStartPos+2]"
                let _29 := calldataload(sum_7)
                /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                let _30 := shr(248, _29)
                let sum_8 := add(var_mainPayloadStartPos, /** @src 0:4235:4236  "4" */ 0x04)
                /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                if gt(var_mainPayloadStartPos, sum_8)
                {
                    mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                    mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x11)
                    revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                }
                /// @src 0:12640:12677  "mainPayloadStartPos + 4 + intByteSize"
                let expr_15 := checked_add_uint256(/** @src 0:12640:12663  "mainPayloadStartPos + 4" */ sum_8, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ shr(248, _25))
                /// @src 0:12711:12743  "tokensStartPos + tokensLength*20"
                let expr_16 := checked_add_uint256(expr_15, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ and(/** @src 0:12728:12743  "tokensLength*20" */ checked_mul_uint8(/** @src 0:12546:12560  "targetByte & 7" */ and(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ _30, /** @src 0:12559:12560  "7" */ 0x07)), /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0xff))
                /// @src 0:12776:12808  "targetsStartPos + swapsLength*20"
                let expr_17 := checked_add_uint256(expr_16, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ and(/** @src 0:12794:12808  "swapsLength*20" */ checked_mul_uint8(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ shr(252, _28)), 0xff))
                /// @src 0:12841:12883  "amountStartPos + amountsLength*intByteSize"
                let expr_18 := checked_add_uint256(expr_17, /** @src 0:12858:12883  "amountsLength*intByteSize" */ checked_mul_uint256(/** @src 0:12135:12150  "targetByte & 15" */ and(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ shr(248, _28), 15), shr(248, _25)))
                if iszero(lt(expr_18, /** @src 0:4071:4079  "msg.data" */ calldatasize()))
                /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                {
                    mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                    mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x32)
                    revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                }
                let sum_9 := add(expr_18, /** @src 0:12038:12039  "1" */ 0x01)
                /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                if gt(expr_18, sum_9)
                {
                    mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                    mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x11)
                    revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                }
                /// @src 0:13208:13239  "uint currPos = zeroForOnesPos+1"
                let var_currPos_2 := sum_9
                /// @src 0:13250:22968  "if (!v3SwapsDone){..."
                switch /** @src 0:13254:13266  "!v3SwapsDone" */ iszero(var_v3SwapsDone)
                case /** @src 0:13250:22968  "if (!v3SwapsDone){..." */ 0 {
                    /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                    let sum_10 := add(var_mainPayloadStartPos, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 3)
                    /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                    if gt(var_mainPayloadStartPos, sum_10)
                    {
                        mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                        mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x11)
                        revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                    }
                    /// @src 0:18806:18828  "revenuePos+intByteSize"
                    let _31 := checked_add_uint256(sum_10, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ shr(248, _25))
                    /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                    if gt(sum_10, _31)
                    {
                        revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                    }
                    /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                    if gt(_31, /** @src 0:4071:4079  "msg.data" */ calldatasize())
                    /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                    {
                        revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                    }
                    /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                    let result := /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ shr(/** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ sub(/** @src 0:11763:11766  "256" */ _27, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ _26), /** @src 0:18778:18830  "bytes32(msg.data[revenuePos:revenuePos+intByteSize])" */ convert_bytes_to_fixedbytes_from_bytes_calldata_to_bytes32(sum_10, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ add(sub(_31, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ var_mainPayloadStartPos), /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ 0xfffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd)))
                    /// @src 0:19033:19857  "if (coinbaseTransfer){..."
                    if /** @src 0:18343:18364  "targetByte & 64 == 64" */ eq(/** @src 0:18343:18358  "targetByte & 64" */ and(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ _30, /** @src 0:18356:18358  "64" */ 0x40), 0x40)
                    /// @src 0:19033:19857  "if (coinbaseTransfer){..."
                    {
                        /// @src 0:19157:19194  "mainPayloadStartPos + 3 + intByteSize"
                        let _32 := checked_add_uint256(/** @src 0:19157:19180  "mainPayloadStartPos + 3" */ sum_10, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ shr(248, _25))
                        if iszero(lt(_32, /** @src 0:4071:4079  "msg.data" */ calldatasize()))
                        /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                        {
                            mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                            mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x32)
                            revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                        }
                        /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                        let r := div(/** @src 0:19492:19533  "revenue * quarterBribePercentageOfRevenue" */ checked_mul_uint256(result, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ shr(248, /** @src 0:19259:19308  "msg.data[QuarterBribePercentageOfRevenueBytesPos]" */ calldataload(_32))), /** @src 0:19536:19539  "250" */ 0xfa)
                        /// @src 0:19661:19683  "WETH.withdraw(revenue)"
                        if iszero(extcodesize(/** @src 0:2484:2526  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2))
                        /// @src 0:19661:19683  "WETH.withdraw(revenue)"
                        {
                            /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                            revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                        }
                        /// @src 0:19661:19683  "WETH.withdraw(revenue)"
                        let _33 := /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ mload(/** @src 0:18356:18358  "64" */ 0x40)
                        /// @src 0:19661:19683  "WETH.withdraw(revenue)"
                        mstore(_33, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ 0x2e1a7d4d00000000000000000000000000000000000000000000000000000000)
                        /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                        mstore(/** @src 0:19661:19683  "WETH.withdraw(revenue)" */ add(_33, /** @src 0:4235:4236  "4" */ 0x04), /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ result)
                        /// @src 0:19661:19683  "WETH.withdraw(revenue)"
                        let _34 := call(gas(), /** @src 0:2484:2526  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2, /** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:19661:19683  "WETH.withdraw(revenue)" */ _33, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 36, /** @src 0:19661:19683  "WETH.withdraw(revenue)" */ _33, /** @src 0:4071:4079  "msg.data" */ _1)
                        /// @src 0:19661:19683  "WETH.withdraw(revenue)"
                        if iszero(_34)
                        {
                            /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                            let pos_3 := /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ mload(/** @src 0:18356:18358  "64" */ 0x40)
                            /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                            returndatacopy(pos_3, /** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ returndatasize())
                            revert(pos_3, returndatasize())
                        }
                        /// @src 0:19661:19683  "WETH.withdraw(revenue)"
                        if _34
                        {
                            finalize_allocation_6340(_33)
                            /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                            if _1
                            {
                                /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                                revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                            }
                        }
                        /// @src 0:19746:19782  "block.coinbase.transfer(bribeAmount)"
                        let _35 := _1
                        if iszero(r) { _35 := 2300 }
                        if iszero(call(_35, /** @src 0:19746:19760  "block.coinbase" */ coinbase(), /** @src 0:19746:19782  "block.coinbase.transfer(bribeAmount)" */ r, /** @src 0:4071:4079  "msg.data" */ _1, _1, _1, _1))
                        /// @src 0:19746:19782  "block.coinbase.transfer(bribeAmount)"
                        {
                            /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                            let pos_4 := /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ mload(/** @src 0:18356:18358  "64" */ 0x40)
                            /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                            returndatacopy(pos_4, /** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ returndatasize())
                            revert(pos_4, returndatasize())
                        }
                    }
                    let sum_11 := add(expr_16, /** @src 0:12741:12743  "20" */ 0x14)
                    /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                    if gt(expr_16, sum_11)
                    {
                        mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                        mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x11)
                        revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                    }
                    /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                    if gt(sum_11, /** @src 0:4071:4079  "msg.data" */ calldatasize())
                    /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                    {
                        revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                    }
                    /// @src 0:20199:20225  "amountStartPos+intByteSize"
                    let _36 := checked_add_uint256(expr_17, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ shr(248, _25))
                    /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                    if gt(expr_17, _36)
                    {
                        revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                    }
                    /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                    if gt(_36, /** @src 0:4071:4079  "msg.data" */ calldatasize())
                    /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                    {
                        revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                    }
                    /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                    let result_1 := /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ shr(/** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ sub(/** @src 0:11763:11766  "256" */ _27, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ _26), /** @src 0:20167:20227  "bytes32(msg.data[amountStartPos:amountStartPos+intByteSize])" */ convert_bytes_to_fixedbytes_from_bytes_calldata_to_bytes32(expr_17, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ sub(_36, expr_17)))
                    /// @src 0:20342:20383  "WETH.transfer(firstTarget, firstAmountIn)"
                    let _37 := /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ mload(/** @src 0:18356:18358  "64" */ 0x40)
                    /// @src 0:20342:20383  "WETH.transfer(firstTarget, firstAmountIn)"
                    mstore(_37, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ 0xa9059cbb00000000000000000000000000000000000000000000000000000000)
                    /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                    mstore(/** @src 0:20342:20383  "WETH.transfer(firstTarget, firstAmountIn)" */ add(_37, /** @src 0:4235:4236  "4" */ 0x04), /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ shr(96, calldataload(expr_16)))
                    mstore(/** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ add(/** @src 0:20342:20383  "WETH.transfer(firstTarget, firstAmountIn)" */ _37, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 36), /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ result_1)
                    /// @src 0:20342:20383  "WETH.transfer(firstTarget, firstAmountIn)"
                    let _38 := call(gas(), /** @src 0:2484:2526  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2, /** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:20342:20383  "WETH.transfer(firstTarget, firstAmountIn)" */ _37, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 68, /** @src 0:20342:20383  "WETH.transfer(firstTarget, firstAmountIn)" */ _37, 32)
                    if iszero(_38)
                    {
                        /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                        let pos_5 := /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ mload(/** @src 0:18356:18358  "64" */ 0x40)
                        /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                        returndatacopy(pos_5, /** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ returndatasize())
                        revert(pos_5, returndatasize())
                    }
                    /// @src 0:20342:20383  "WETH.transfer(firstTarget, firstAmountIn)"
                    let expr_19 := _1
                    if _38
                    {
                        let _39 := 32
                        if gt(_39, returndatasize()) { _39 := returndatasize() }
                        finalize_allocation(_37, _39)
                        /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                        if slt(sub(/** @src 0:20342:20383  "WETH.transfer(firstTarget, firstAmountIn)" */ add(_37, _39), /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ _37), /** @src 0:20342:20383  "WETH.transfer(firstTarget, firstAmountIn)" */ 32)
                        /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                        {
                            /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                            revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                        }
                        /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                        let value_1 := mload(_37)
                        if iszero(eq(value_1, iszero(iszero(value_1))))
                        {
                            revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                        }
                        /// @src 0:20342:20383  "WETH.transfer(firstTarget, firstAmountIn)"
                        expr_19 := value_1
                    }
                    /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                    if iszero(expr_19)
                    {
                        let memPtr := /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ mload(/** @src 0:18356:18358  "64" */ 0x40)
                        /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                        mstore(memPtr, 0x08c379a000000000000000000000000000000000000000000000000000000000)
                        mstore(add(memPtr, /** @src 0:4235:4236  "4" */ 0x04), /** @src 0:20342:20383  "WETH.transfer(firstTarget, firstAmountIn)" */ 32)
                        /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                        mstore(add(memPtr, 36), 31)
                        mstore(add(memPtr, 68), "transfer to first target failed")
                        revert(memPtr, 100)
                    }
                    /// @src 0:20843:20854  "uint8 i = 0"
                    let var_i := _1
                    /// @src 0:20838:22869  "for (uint8 i = 0; i < v2SwapDataLength; ++i) {..."
                    for { }
                    /** @src 0:20856:20876  "i < v2SwapDataLength" */ lt(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ and(/** @src 0:20856:20876  "i < v2SwapDataLength" */ var_i, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0xff), and(shr(251, _29), /** @src 0:12559:12560  "7" */ 0x07))
                    /// @src 0:20843:20854  "uint8 i = 0"
                    {
                        /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                        if eq(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ and(/** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ var_i, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0xff), 0xff)
                        /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                        {
                            /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                            mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                            mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x11)
                            revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 36)
                        }
                        /// @src 0:20878:20881  "++i"
                        var_i := /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ add(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ and(/** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ var_i, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0xff), /** @src 0:12038:12039  "1" */ 0x01)
                    }
                    /// @src 0:20878:20881  "++i"
                    {
                        /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                        if iszero(lt(var_currPos_2, /** @src 0:4071:4079  "msg.data" */ calldatasize()))
                        /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                        {
                            mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                            mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x32)
                            revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 36)
                        }
                        /// @src 0:21048:21065  "msg.data[currPos]"
                        let _40 := calldataload(var_currPos_2)
                        /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                        let result_2 := shr(251, _40)
                        /// @src 0:21170:21184  "targetByte & 7"
                        let _41 := and(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ shr(248, _40), /** @src 0:12559:12560  "7" */ 0x07)
                        /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                        if eq(var_currPos_2, 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff)
                        {
                            /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                            mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                            mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x11)
                            revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 36)
                        }
                        /// @src 0:21219:21229  "currPos ++"
                        var_currPos_2 := /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ add(var_currPos_2, /** @src 0:12038:12039  "1" */ 0x01)
                        /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                        if iszero(or(iszero(result_2), eq(/** @src 0:12741:12743  "20" */ 0x14, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ div(mul(result_2, /** @src 0:12741:12743  "20" */ 0x14), /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ result_2))))
                        {
                            /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                            mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                            mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x11)
                            revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 36)
                        }
                        /// @src 0:21476:21506  "targetsStartPos + swapIndex*20"
                        let expr_20 := checked_add_uint256(expr_16, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ mul(result_2, /** @src 0:12741:12743  "20" */ 0x14))
                        /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                        let sum_12 := add(result_2, /** @src 0:12038:12039  "1" */ 0x01)
                        /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                        if gt(result_2, sum_12)
                        {
                            mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                            mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x11)
                            revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 36)
                        }
                        let product_1 := mul(sum_12, /** @src 0:12741:12743  "20" */ 0x14)
                        /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                        if iszero(or(iszero(sum_12), eq(/** @src 0:12741:12743  "20" */ 0x14, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ div(product_1, sum_12))))
                        {
                            /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                            mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                            mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x11)
                            revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 36)
                        }
                        /// @src 0:21507:21541  "targetsStartPos + (swapIndex+1)*20"
                        let _42 := checked_add_uint256(expr_16, /** @src 0:21525:21541  "(swapIndex+1)*20" */ product_1)
                        /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                        if gt(expr_20, _42)
                        {
                            revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                        }
                        /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                        if gt(_42, /** @src 0:4071:4079  "msg.data" */ calldatasize())
                        /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                        {
                            revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                        }
                        /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                        let _43 := shr(96, /** @src 0:21459:21543  "bytes20(msg.data[targetsStartPos + swapIndex*20:targetsStartPos + (swapIndex+1)*20])" */ convert_bytes_to_fixedbytes_from_bytes_calldata_to_bytes20(expr_20, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ sub(_42, expr_20)))
                        /// @src 0:21609:21643  "targetsStartPos + (swapIndex+1)*20"
                        let expr_21 := checked_add_uint256(expr_16, /** @src 0:21627:21643  "(swapIndex+1)*20" */ product_1)
                        /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                        let sum_13 := add(result_2, /** @src 0:12511:12512  "2" */ 0x02)
                        /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                        if gt(result_2, sum_13)
                        {
                            mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                            mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x11)
                            revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 36)
                        }
                        let product_2 := mul(sum_13, /** @src 0:12741:12743  "20" */ 0x14)
                        /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                        if iszero(or(iszero(sum_13), eq(/** @src 0:12741:12743  "20" */ 0x14, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ div(product_2, sum_13))))
                        {
                            /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                            mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                            mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x11)
                            revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 36)
                        }
                        /// @src 0:21644:21678  "targetsStartPos + (swapIndex+2)*20"
                        let _44 := checked_add_uint256(expr_16, /** @src 0:21662:21678  "(swapIndex+2)*20" */ product_2)
                        /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                        if gt(expr_21, _44)
                        {
                            revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                        }
                        /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                        if gt(_44, /** @src 0:4071:4079  "msg.data" */ calldatasize())
                        /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                        {
                            revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                        }
                        /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                        let _45 := shr(96, /** @src 0:21592:21680  "bytes20(msg.data[targetsStartPos + (swapIndex+1)*20:targetsStartPos + (swapIndex+2)*20])" */ convert_bytes_to_fixedbytes_from_bytes_calldata_to_bytes20(expr_21, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ sub(_44, expr_21)))
                        /// @src 0:21739:21779  "amountStartPos + amountIndex*intByteSize"
                        let expr_22 := checked_add_uint256(expr_17, /** @src 0:21756:21779  "amountIndex*intByteSize" */ checked_mul_uint256(_41, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ shr(248, _25)))
                        let sum_14 := add(_41, /** @src 0:12038:12039  "1" */ 0x01)
                        /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                        if gt(_41, sum_14)
                        {
                            mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                            mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x11)
                            revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 36)
                        }
                        /// @src 0:21780:21824  "amountStartPos + (amountIndex+1)*intByteSize"
                        let _46 := checked_add_uint256(expr_17, /** @src 0:21797:21824  "(amountIndex+1)*intByteSize" */ checked_mul_uint256(/** @src 0:21798:21811  "amountIndex+1" */ sum_14, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ shr(248, _25)))
                        /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                        if gt(expr_22, _46)
                        {
                            revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                        }
                        /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                        if gt(_46, /** @src 0:4071:4079  "msg.data" */ calldatasize())
                        /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                        {
                            revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                        }
                        /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                        let result_3 := /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ shr(/** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ sub(/** @src 0:11763:11766  "256" */ _27, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ _26), /** @src 0:21722:21826  "bytes32(msg.data[amountStartPos + amountIndex*intByteSize:amountStartPos + (amountIndex+1)*intByteSize])" */ convert_bytes_to_fixedbytes_from_bytes_calldata_to_bytes32(expr_22, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ sub(_46, expr_22)))
                        /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                        let result_4 := /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ shl(/** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ result_2, /** @src 0:12038:12039  "1" */ 0x01)
                        /// @src 0:22612:22854  "if (zeroForOne) {..."
                        switch /** @src 0:21944:21972  "(zeroForOnes & mask) == mask" */ eq(/** @src 0:21945:21963  "zeroForOnes & mask" */ and(and(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ shr(248, /** @src 0:13113:13137  "msg.data[zeroForOnesPos]" */ calldataload(expr_18)), /** @src 0:21945:21963  "zeroForOnes & mask" */ result_4), /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0xff), /** @src 0:21944:21972  "(zeroForOnes & mask) == mask" */ result_4)
                        case /** @src 0:22612:22854  "if (zeroForOne) {..." */ 0 {
                            /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                            let memPtr_1 := /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ mload(/** @src 0:18356:18358  "64" */ 0x40)
                            /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                            finalize_allocation_8900(memPtr_1)
                            mstore(memPtr_1, /** @src 0:4071:4079  "msg.data" */ _1)
                            /// @src 0:22767:22834  "IUniswapV2Pair(target).swap(amountOut, 0, nextTarget, new bytes(0))"
                            if iszero(extcodesize(_43))
                            {
                                /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                                revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                            }
                            /// @src 0:22767:22834  "IUniswapV2Pair(target).swap(amountOut, 0, nextTarget, new bytes(0))"
                            let _47 := /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ mload(/** @src 0:18356:18358  "64" */ 0x40)
                            /// @src 0:22767:22834  "IUniswapV2Pair(target).swap(amountOut, 0, nextTarget, new bytes(0))"
                            mstore(_47, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ 0x022c0d9f00000000000000000000000000000000000000000000000000000000)
                            /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                            mstore(/** @src 0:22767:22834  "IUniswapV2Pair(target).swap(amountOut, 0, nextTarget, new bytes(0))" */ add(_47, /** @src 0:4235:4236  "4" */ 0x04), /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ result_3)
                            mstore(/** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ add(/** @src 0:22767:22834  "IUniswapV2Pair(target).swap(amountOut, 0, nextTarget, new bytes(0))" */ _47, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 36), /** @src 0:4071:4079  "msg.data" */ _1)
                            /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                            mstore(/** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ add(/** @src 0:22767:22834  "IUniswapV2Pair(target).swap(amountOut, 0, nextTarget, new bytes(0))" */ _47, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 68), /** @src 0:3159:3169  "0x3ccfd60b" */ _45)
                            /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                            mstore(add(/** @src 0:22767:22834  "IUniswapV2Pair(target).swap(amountOut, 0, nextTarget, new bytes(0))" */ _47, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 100), 128)
                            /// @src 0:22767:22834  "IUniswapV2Pair(target).swap(amountOut, 0, nextTarget, new bytes(0))"
                            let _48 := call(gas(), _43, /** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:22767:22834  "IUniswapV2Pair(target).swap(amountOut, 0, nextTarget, new bytes(0))" */ _47, sub(/** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ abi_encode_bytes(memPtr_1, add(/** @src 0:22767:22834  "IUniswapV2Pair(target).swap(amountOut, 0, nextTarget, new bytes(0))" */ _47, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 132)), /** @src 0:22767:22834  "IUniswapV2Pair(target).swap(amountOut, 0, nextTarget, new bytes(0))" */ _47), _47, /** @src 0:4071:4079  "msg.data" */ _1)
                            /// @src 0:22767:22834  "IUniswapV2Pair(target).swap(amountOut, 0, nextTarget, new bytes(0))"
                            if iszero(_48)
                            {
                                /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                                let pos_6 := /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ mload(/** @src 0:18356:18358  "64" */ 0x40)
                                /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                                returndatacopy(pos_6, /** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ returndatasize())
                                revert(pos_6, returndatasize())
                            }
                            /// @src 0:22767:22834  "IUniswapV2Pair(target).swap(amountOut, 0, nextTarget, new bytes(0))"
                            if _48
                            {
                                finalize_allocation_6340(_47)
                                /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                                if _1
                                {
                                    /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                                    revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                                }
                            }
                        }
                        default /// @src 0:22612:22854  "if (zeroForOne) {..."
                        {
                            /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                            let memPtr_2 := /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ mload(/** @src 0:18356:18358  "64" */ 0x40)
                            /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                            finalize_allocation_8900(memPtr_2)
                            mstore(memPtr_2, /** @src 0:4071:4079  "msg.data" */ _1)
                            /// @src 0:22651:22718  "IUniswapV2Pair(target).swap(0, amountOut, nextTarget, new bytes(0))"
                            if iszero(extcodesize(_43))
                            {
                                /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                                revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                            }
                            /// @src 0:22651:22718  "IUniswapV2Pair(target).swap(0, amountOut, nextTarget, new bytes(0))"
                            let _49 := /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ mload(/** @src 0:18356:18358  "64" */ 0x40)
                            /// @src 0:22651:22718  "IUniswapV2Pair(target).swap(0, amountOut, nextTarget, new bytes(0))"
                            mstore(_49, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ 0x022c0d9f00000000000000000000000000000000000000000000000000000000)
                            /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                            mstore(/** @src 0:22651:22718  "IUniswapV2Pair(target).swap(0, amountOut, nextTarget, new bytes(0))" */ add(_49, /** @src 0:4235:4236  "4" */ 0x04), /** @src 0:4071:4079  "msg.data" */ _1)
                            /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                            mstore(/** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ add(/** @src 0:22651:22718  "IUniswapV2Pair(target).swap(0, amountOut, nextTarget, new bytes(0))" */ _49, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 36), /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ result_3)
                            mstore(/** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ add(/** @src 0:22651:22718  "IUniswapV2Pair(target).swap(0, amountOut, nextTarget, new bytes(0))" */ _49, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 68), /** @src 0:3159:3169  "0x3ccfd60b" */ _45)
                            /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                            mstore(add(/** @src 0:22651:22718  "IUniswapV2Pair(target).swap(0, amountOut, nextTarget, new bytes(0))" */ _49, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 100), 128)
                            /// @src 0:22651:22718  "IUniswapV2Pair(target).swap(0, amountOut, nextTarget, new bytes(0))"
                            let _50 := call(gas(), _43, /** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:22651:22718  "IUniswapV2Pair(target).swap(0, amountOut, nextTarget, new bytes(0))" */ _49, sub(/** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ abi_encode_bytes(memPtr_2, add(/** @src 0:22651:22718  "IUniswapV2Pair(target).swap(0, amountOut, nextTarget, new bytes(0))" */ _49, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 132)), /** @src 0:22651:22718  "IUniswapV2Pair(target).swap(0, amountOut, nextTarget, new bytes(0))" */ _49), _49, /** @src 0:4071:4079  "msg.data" */ _1)
                            /// @src 0:22651:22718  "IUniswapV2Pair(target).swap(0, amountOut, nextTarget, new bytes(0))"
                            if iszero(_50)
                            {
                                /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                                let pos_7 := /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ mload(/** @src 0:18356:18358  "64" */ 0x40)
                                /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                                returndatacopy(pos_7, /** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ returndatasize())
                                revert(pos_7, returndatasize())
                            }
                            /// @src 0:22651:22718  "IUniswapV2Pair(target).swap(0, amountOut, nextTarget, new bytes(0))"
                            if _50
                            {
                                finalize_allocation_6340(_49)
                                /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                                if _1
                                {
                                    /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                                    revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                                }
                            }
                        }
                    }
                    /// @src 0:22935:22942  "return;"
                    leave
                }
                default /// @src 0:13250:22968  "if (!v3SwapsDone){..."
                {
                    /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                    if iszero(lt(var_currPos_2, /** @src 0:4071:4079  "msg.data" */ calldatasize()))
                    /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                    {
                        mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                        mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x32)
                        revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                    }
                    /// @src 0:14126:14143  "msg.data[currPos]"
                    let _51 := calldataload(var_currPos_2)
                    /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                    let _52 := and(shr(251, _51), /** @src 0:12559:12560  "7" */ 0x07)
                    /// @src 0:14249:14263  "targetByte & 7"
                    let expr_23 := and(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ shr(248, _51), /** @src 0:12559:12560  "7" */ 0x07)
                    /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                    let sum_15 := add(var_currPos_2, /** @src 0:12038:12039  "1" */ 0x01)
                    /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                    if gt(var_currPos_2, sum_15)
                    {
                        mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                        mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x11)
                        revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                    }
                    if iszero(lt(sum_15, /** @src 0:4071:4079  "msg.data" */ calldatasize()))
                    /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                    {
                        mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                        mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x32)
                        revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                    }
                    /// @src 0:15140:15170  "targetsStartPos + swapIndex*20"
                    let expr_24 := checked_add_uint256(expr_16, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ and(/** @src 0:15158:15170  "swapIndex*20" */ checked_mul_uint8(_52), /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0xff))
                    /// @src 0:15171:15205  "targetsStartPos + (swapIndex+1)*20"
                    let _53 := checked_add_uint256(expr_16, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ and(/** @src 0:15189:15205  "(swapIndex+1)*20" */ checked_mul_uint8(/** @src 0:15190:15201  "swapIndex+1" */ checked_add_uint8(_52)), /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0xff))
                    /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                    if gt(expr_24, _53)
                    {
                        revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                    }
                    /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                    if gt(_53, /** @src 0:4071:4079  "msg.data" */ calldatasize())
                    /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                    {
                        revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                    }
                    /// @src 0:15123:15207  "bytes20(msg.data[targetsStartPos + swapIndex*20:targetsStartPos + (swapIndex+1)*20])"
                    let _54 := convert_bytes_to_fixedbytes_from_bytes_calldata_to_bytes20(expr_24, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ sub(_53, expr_24))
                    /// @src 0:15223:15241  "address nextTarget"
                    let var_nextTarget := _1
                    /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                    let diff := add(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ shr(252, _28), /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff)
                    if gt(diff, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0xff)
                    /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                    {
                        /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                        mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                        mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x11)
                        revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                    }
                    /// @src 0:15256:15550  "if (swapIndex == swapsLength-1){//if last address, set next target to this contract..."
                    switch /** @src 0:15260:15286  "swapIndex == swapsLength-1" */ eq(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ _52, and(diff, 0xff))
                    case /** @src 0:15256:15550  "if (swapIndex == swapsLength-1){//if last address, set next target to this contract..." */ 0 {
                        /// @src 0:15462:15496  "targetsStartPos + (swapIndex+1)*20"
                        let expr_25 := checked_add_uint256(expr_16, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ and(/** @src 0:15480:15496  "(swapIndex+1)*20" */ checked_mul_uint8(/** @src 0:15481:15492  "swapIndex+1" */ checked_add_uint8(_52)), /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0xff))
                        /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                        let sum_16 := add(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ _52, /** @src 0:12511:12512  "2" */ 0x02)
                        /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                        if gt(sum_16, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0xff)
                        /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                        {
                            /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                            mstore(/** @src 0:4071:4079  "msg.data" */ 0, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                            mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x11)
                            revert(/** @src 0:4071:4079  "msg.data" */ 0, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                        }
                        /// @src 0:15497:15531  "targetsStartPos + (swapIndex+2)*20"
                        let _55 := checked_add_uint256(expr_16, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ and(/** @src 0:15515:15531  "(swapIndex+2)*20" */ checked_mul_uint8(/** @src 0:15516:15527  "swapIndex+2" */ sum_16), /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0xff))
                        /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                        if gt(expr_25, _55)
                        {
                            revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                        }
                        /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                        if gt(_55, /** @src 0:4071:4079  "msg.data" */ calldatasize())
                        /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                        {
                            revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                        }
                        /// @src 0:15424:15534  "nextTarget = address(bytes20(msg.data[targetsStartPos + (swapIndex+1)*20:targetsStartPos + (swapIndex+2)*20]))"
                        var_nextTarget := /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ shr(96, /** @src 0:15445:15533  "bytes20(msg.data[targetsStartPos + (swapIndex+1)*20:targetsStartPos + (swapIndex+2)*20])" */ convert_bytes_to_fixedbytes_from_bytes_calldata_to_bytes20(expr_25, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ sub(_55, expr_25)))
                    }
                    default /// @src 0:15256:15550  "if (swapIndex == swapsLength-1){//if last address, set next target to this contract..."
                    {
                        /// @src 0:15357:15383  "nextTarget = address(this)"
                        var_nextTarget := /** @src 0:15378:15382  "this" */ address()
                    }
                    /// @src 0:15743:15783  "amountStartPos + amountIndex*intByteSize"
                    let expr_26 := checked_add_uint256(expr_17, /** @src 0:15760:15783  "amountIndex*intByteSize" */ checked_mul_uint256(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ expr_23, shr(248, _25)))
                    /// @src 0:15784:15828  "amountStartPos + (amountIndex+1)*intByteSize"
                    let _56 := checked_add_uint256(expr_17, /** @src 0:15801:15828  "(amountIndex+1)*intByteSize" */ checked_mul_uint256(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ and(/** @src 0:15802:15815  "amountIndex+1" */ checked_add_uint8(expr_23), /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0xff), shr(248, _25)))
                    /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                    if gt(expr_26, _56)
                    {
                        revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                    }
                    /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                    if gt(_56, /** @src 0:4071:4079  "msg.data" */ calldatasize())
                    /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                    {
                        revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                    }
                    /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                    let result_5 := /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ shr(/** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ sub(/** @src 0:11763:11766  "256" */ _27, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ _26), /** @src 0:15726:15830  "bytes32(msg.data[amountStartPos + amountIndex*intByteSize:amountStartPos + (amountIndex+1)*intByteSize])" */ convert_bytes_to_fixedbytes_from_bytes_calldata_to_bytes32(expr_26, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ sub(_56, expr_26)))
                    /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                    let result_6 := /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ shl(/** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ _52, /** @src 0:12038:12039  "1" */ 0x01)
                    /// @src 0:16029:16057  "(zeroForOnes & mask) == mask"
                    let expr_27 := eq(/** @src 0:16030:16048  "zeroForOnes & mask" */ and(and(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ shr(248, /** @src 0:13113:13137  "msg.data[zeroForOnesPos]" */ calldataload(expr_18)), /** @src 0:16030:16048  "zeroForOnes & mask" */ result_6), /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0xff), /** @src 0:16029:16057  "(zeroForOnes & mask) == mask" */ result_6)
                    /// @src 0:16102:16161  "zeroForOne ? MIN_SQRT_RATIO_PLUS_1 : MAX_SQRT_RATIO_MINUS_1"
                    let expr_28 := _1
                    switch expr_27
                    case 0 {
                        expr_28 := /** @src 0:3540:3589  "1461446703485210103287273052203988822378723970341" */ 0xfffd8963efd1fc6a506488495d951d5263988d25
                    }
                    default /// @src 0:16102:16161  "zeroForOne ? MIN_SQRT_RATIO_PLUS_1 : MAX_SQRT_RATIO_MINUS_1"
                    {
                        expr_28 := /** @src 0:3350:3360  "4295128740" */ 0x01000276a4
                    }
                    /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                    let sum_17 := add(var_currPos_2, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 3)
                    /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                    if gt(var_currPos_2, sum_17)
                    {
                        mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                        mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x11)
                        revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                    }
                    /// @src 0:16208:16254  "uint v3SwapDataEndPos = v3SwapDataStartPos + 3"
                    let var_v3SwapDataEndPos := sum_17
                    /// @src 0:16269:16416  "if(factoryIndex == 1){// if factory is not indexed; value of 1 means custom factory address..."
                    if /** @src 0:16272:16289  "factoryIndex == 1" */ eq(/** @src 0:14729:14743  "targetByte & 7" */ and(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ shr(248, /** @src 0:14675:14692  "msg.data[currPos]" */ calldataload(sum_15)), /** @src 0:12559:12560  "7" */ 0x07), /** @src 0:12038:12039  "1" */ 0x01)
                    /// @src 0:16269:16416  "if(factoryIndex == 1){// if factory is not indexed; value of 1 means custom factory address..."
                    {
                        /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                        let sum_18 := add(var_currPos_2, 23)
                        if gt(sum_17, sum_18)
                        {
                            mstore(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 35408467139433450592217433187231851964531694900788300625387963629091585785856)
                            mstore(/** @src 0:4235:4236  "4" */ 0x04, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x11)
                            revert(/** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0x24)
                        }
                        /// @src 0:16378:16400  "v3SwapDataEndPos += 20"
                        var_v3SwapDataEndPos := sum_18
                    }
                    /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                    if gt(var_currPos_2, var_v3SwapDataEndPos)
                    {
                        revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                    }
                    /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                    if gt(var_v3SwapDataEndPos, /** @src 0:4071:4079  "msg.data" */ calldatasize())
                    /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                    {
                        revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                    }
                    /// @src 0:16493:16571  "bytes memory currentV3SwapData = msg.data[v3SwapDataStartPos:v3SwapDataEndPos]"
                    let var_currentV3SwapData_mpos := /** @src 0:3350:3360  "4295128740" */ abi_decode_available_length_bytes(/** @src 0:16493:16571  "bytes memory currentV3SwapData = msg.data[v3SwapDataStartPos:v3SwapDataEndPos]" */ var_currPos_2, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ sub(var_v3SwapDataEndPos, var_currPos_2), /** @src 0:4071:4079  "msg.data" */ calldatasize())
                    /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                    if gt(var_mainPayloadStartPos, var_currPos_2)
                    {
                        revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                    }
                    /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                    if gt(var_currPos_2, /** @src 0:4071:4079  "msg.data" */ calldatasize())
                    /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                    {
                        revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                    }
                    /// @src 0:16914:16992  "bytes memory newPayloadHead = msg.data[mainPayloadStartPos:v3SwapDataStartPos]"
                    let var_newPayloadHead_mpos := /** @src 0:3350:3360  "4295128740" */ abi_decode_available_length_bytes(/** @src 0:16914:16992  "bytes memory newPayloadHead = msg.data[mainPayloadStartPos:v3SwapDataStartPos]" */ var_mainPayloadStartPos, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ sub(var_currPos_2, var_mainPayloadStartPos), /** @src 0:4071:4079  "msg.data" */ calldatasize())
                    /// @src 0:17007:17064  "bytes memory newPayloadTail = msg.data[v3SwapDataEndPos:]"
                    let var_newPayloadTail_mpos := /** @src 0:3350:3360  "4295128740" */ abi_decode_available_length_bytes(/** @src 0:17007:17064  "bytes memory newPayloadTail = msg.data[v3SwapDataEndPos:]" */ var_v3SwapDataEndPos, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ sub(/** @src 0:4071:4079  "msg.data" */ calldatasize(), /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ var_v3SwapDataEndPos), /** @src 0:4071:4079  "msg.data" */ calldatasize())
                    /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                    let _57 := 64
                    /// @src 0:17105:17172  "abi.encodePacked(currentV3SwapData, newPayloadHead, newPayloadTail)"
                    let expr_1041_mpos := /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ mload(_57)
                    /// @src 0:17105:17172  "abi.encodePacked(currentV3SwapData, newPayloadHead, newPayloadTail)"
                    let _58 := 0x20
                    /// @src 0:3350:3360  "4295128740"
                    let length := /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ mload(/** @src 0:3350:3360  "4295128740" */ var_currentV3SwapData_mpos)
                    copy_memory_to_memory_with_cleanup(add(var_currentV3SwapData_mpos, /** @src 0:17105:17172  "abi.encodePacked(currentV3SwapData, newPayloadHead, newPayloadTail)" */ _58), add(expr_1041_mpos, _58), /** @src 0:3350:3360  "4295128740" */ length)
                    let _59 := add(/** @src 0:17105:17172  "abi.encodePacked(currentV3SwapData, newPayloadHead, newPayloadTail)" */ expr_1041_mpos, /** @src 0:3350:3360  "4295128740" */ length)
                    let length_1 := /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ mload(/** @src 0:3350:3360  "4295128740" */ var_newPayloadHead_mpos)
                    copy_memory_to_memory_with_cleanup(add(var_newPayloadHead_mpos, /** @src 0:17105:17172  "abi.encodePacked(currentV3SwapData, newPayloadHead, newPayloadTail)" */ _58), /** @src 0:3350:3360  "4295128740" */ add(_59, /** @src 0:17105:17172  "abi.encodePacked(currentV3SwapData, newPayloadHead, newPayloadTail)" */ _58), /** @src 0:3350:3360  "4295128740" */ length_1)
                    let _60 := add(_59, length_1)
                    let length_2 := /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ mload(/** @src 0:3350:3360  "4295128740" */ var_newPayloadTail_mpos)
                    copy_memory_to_memory_with_cleanup(add(var_newPayloadTail_mpos, /** @src 0:17105:17172  "abi.encodePacked(currentV3SwapData, newPayloadHead, newPayloadTail)" */ _58), /** @src 0:3350:3360  "4295128740" */ add(_60, /** @src 0:17105:17172  "abi.encodePacked(currentV3SwapData, newPayloadHead, newPayloadTail)" */ _58), /** @src 0:3350:3360  "4295128740" */ length_2)
                    /// @src 0:17105:17172  "abi.encodePacked(currentV3SwapData, newPayloadHead, newPayloadTail)"
                    let _61 := sub(/** @src 0:3350:3360  "4295128740" */ add(_60, length_2), /** @src 0:17105:17172  "abi.encodePacked(currentV3SwapData, newPayloadHead, newPayloadTail)" */ expr_1041_mpos)
                    mstore(expr_1041_mpos, _61)
                    finalize_allocation(expr_1041_mpos, add(_61, _58))
                    /// @src 0:17446:17544  "IUniswapV3Pool(target).swap(nextTarget, zeroForOne, int256(amount), sqrtPriceLimitX96, newPayload)"
                    let _62 := /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ mload(_57)
                    /// @src 0:17446:17544  "IUniswapV3Pool(target).swap(nextTarget, zeroForOne, int256(amount), sqrtPriceLimitX96, newPayload)"
                    mstore(_62, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ 0x128acb0800000000000000000000000000000000000000000000000000000000)
                    /// @src 0:3159:3169  "0x3ccfd60b"
                    let _63 := 0xffffffffffffffffffffffffffffffffffffffff
                    /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                    mstore(/** @src 0:17446:17544  "IUniswapV3Pool(target).swap(nextTarget, zeroForOne, int256(amount), sqrtPriceLimitX96, newPayload)" */ add(_62, /** @src 0:4235:4236  "4" */ 0x04), /** @src 0:3159:3169  "0x3ccfd60b" */ and(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ var_nextTarget, /** @src 0:3159:3169  "0x3ccfd60b" */ _63))
                    /// @src 0:3350:3360  "4295128740"
                    mstore(add(/** @src 0:17446:17544  "IUniswapV3Pool(target).swap(nextTarget, zeroForOne, int256(amount), sqrtPriceLimitX96, newPayload)" */ _62, /** @src 0:3350:3360  "4295128740" */ 36), /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ expr_27)
                    /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                    mstore(/** @src 0:3350:3360  "4295128740" */ add(/** @src 0:17446:17544  "IUniswapV3Pool(target).swap(nextTarget, zeroForOne, int256(amount), sqrtPriceLimitX96, newPayload)" */ _62, /** @src 0:3350:3360  "4295128740" */ 68), /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ result_5)
                    mstore(/** @src 0:3350:3360  "4295128740" */ add(/** @src 0:17446:17544  "IUniswapV3Pool(target).swap(nextTarget, zeroForOne, int256(amount), sqrtPriceLimitX96, newPayload)" */ _62, /** @src 0:3350:3360  "4295128740" */ 100), /** @src 0:3159:3169  "0x3ccfd60b" */ and(/** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ expr_28, /** @src 0:3159:3169  "0x3ccfd60b" */ _63))
                    /// @src 0:3350:3360  "4295128740"
                    mstore(add(/** @src 0:17446:17544  "IUniswapV3Pool(target).swap(nextTarget, zeroForOne, int256(amount), sqrtPriceLimitX96, newPayload)" */ _62, /** @src 0:3350:3360  "4295128740" */ 132), 160)
                    /// @src 0:17446:17544  "IUniswapV3Pool(target).swap(nextTarget, zeroForOne, int256(amount), sqrtPriceLimitX96, newPayload)"
                    let _64 := call(gas(), /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ shr(96, /** @src 0:15123:15207  "bytes20(msg.data[targetsStartPos + swapIndex*20:targetsStartPos + (swapIndex+1)*20])" */ _54), /** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:17446:17544  "IUniswapV3Pool(target).swap(nextTarget, zeroForOne, int256(amount), sqrtPriceLimitX96, newPayload)" */ _62, sub(/** @src 0:3350:3360  "4295128740" */ abi_encode_bytes(expr_1041_mpos, add(/** @src 0:17446:17544  "IUniswapV3Pool(target).swap(nextTarget, zeroForOne, int256(amount), sqrtPriceLimitX96, newPayload)" */ _62, /** @src 0:3350:3360  "4295128740" */ 164)), /** @src 0:17446:17544  "IUniswapV3Pool(target).swap(nextTarget, zeroForOne, int256(amount), sqrtPriceLimitX96, newPayload)" */ _62), _62, /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ _57)
                    /// @src 0:17446:17544  "IUniswapV3Pool(target).swap(nextTarget, zeroForOne, int256(amount), sqrtPriceLimitX96, newPayload)"
                    if iszero(_64)
                    {
                        /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                        let pos_8 := /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ mload(_57)
                        /// @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)"
                        returndatacopy(pos_8, /** @src 0:4071:4079  "msg.data" */ _1, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ returndatasize())
                        revert(pos_8, returndatasize())
                    }
                    /// @src 0:17446:17544  "IUniswapV3Pool(target).swap(nextTarget, zeroForOne, int256(amount), sqrtPriceLimitX96, newPayload)"
                    if _64
                    {
                        let _65 := /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ _57
                        /// @src 0:17446:17544  "IUniswapV3Pool(target).swap(nextTarget, zeroForOne, int256(amount), sqrtPriceLimitX96, newPayload)"
                        if gt(/** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ _57, /** @src 0:17446:17544  "IUniswapV3Pool(target).swap(nextTarget, zeroForOne, int256(amount), sqrtPriceLimitX96, newPayload)" */ returndatasize()) { _65 := returndatasize() }
                        finalize_allocation(_62, _65)
                        /// @src 0:3350:3360  "4295128740"
                        if slt(sub(/** @src 0:17446:17544  "IUniswapV3Pool(target).swap(nextTarget, zeroForOne, int256(amount), sqrtPriceLimitX96, newPayload)" */ add(_62, _65), /** @src 0:3350:3360  "4295128740" */ _62), /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ _57)
                        /// @src 0:3350:3360  "4295128740"
                        {
                            /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                            revert(/** @src 0:4071:4079  "msg.data" */ _1, _1)
                        }
                    }
                    /// @src 0:17814:17821  "return;"
                    leave
                }
            }
            /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
            function abi_encode_stringliteral_b57e(headStart) -> tail
            {
                mstore(headStart, 32)
                /// @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
                mstore(/** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ add(headStart, 32), 17)
                mstore(/** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ add(/** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ headStart, /** @src 0:2667:2709  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" */ 64), /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ "Invalid fee index")
                tail := add(headStart, 96)
            }
            /// @ast-id 1519 @src 0:24730:25153  "function getFee(uint8 feeIndex) internal pure returns (uint24) {..."
            function fun_getFee(var_feeIndex) -> var
            {
                /// @src 0:24785:24791  "uint24"
                var := /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ 0
                /// @src 0:24810:24822  "feeIndex < 4"
                let _1 := /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ and(/** @src 0:24810:24822  "feeIndex < 4" */ var_feeIndex, /** @src 0:2478:2527  "IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)" */ 0xff)
                /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                if iszero(/** @src 0:24810:24822  "feeIndex < 4" */ lt(_1, /** @src 0:24821:24822  "4" */ 0x04))
                /// @src 0:2137:25520  "contract UniswapMixedExecutor {..."
                {
                    let memPtr := mload(64)
                    mstore(memPtr, 0x08c379a000000000000000000000000000000000000000000000000000000000)
                    revert(memPtr, sub(abi_encode_stringliteral_b57e(add(memPtr, /** @src 0:24821:24822  "4" */ 0x04)), /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ memPtr))
                }
                /// @src 0:24861:25146  "if (feeIndex == 0) {..."
                switch /** @src 0:24865:24878  "feeIndex == 0" */ iszero(_1)
                case /** @src 0:24861:25146  "if (feeIndex == 0) {..." */ 0 {
                    /// @src 0:24919:25146  "if (feeIndex == 1) {..."
                    switch /** @src 0:24923:24936  "feeIndex == 1" */ eq(_1, /** @src 0:24935:24936  "1" */ 0x01)
                    case /** @src 0:24919:25146  "if (feeIndex == 1) {..." */ 0 {
                        /// @src 0:24978:25146  "if (feeIndex == 2) {..."
                        switch /** @src 0:24982:24995  "feeIndex == 2" */ eq(_1, /** @src 0:24994:24995  "2" */ 0x02)
                        case /** @src 0:24978:25146  "if (feeIndex == 2) {..." */ 0 {
                            /// @src 0:25038:25146  "if (feeIndex == 3) {..."
                            switch /** @src 0:25042:25055  "feeIndex == 3" */ eq(_1, /** @src 0:25054:25055  "3" */ 0x03)
                            case /** @src 0:25038:25146  "if (feeIndex == 3) {..." */ 0 {
                                /// @src 0:25109:25136  "revert(\"Invalid fee index\")"
                                let _2 := /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ mload(64)
                                /// @src 0:25109:25136  "revert(\"Invalid fee index\")"
                                mstore(_2, 3963877391197344453575983046348115674221700746820753546331534351508065746944)
                                revert(_2, sub(abi_encode_stringliteral_b57e(add(_2, /** @src 0:24821:24822  "4" */ 0x04)), /** @src 0:25109:25136  "revert(\"Invalid fee index\")" */ _2))
                            }
                            default /// @src 0:25038:25146  "if (feeIndex == 3) {..."
                            {
                                /// @src 0:25070:25080  "return 100"
                                var := /** @src 0:25077:25080  "100" */ 0x64
                                /// @src 0:25070:25080  "return 100"
                                leave
                            }
                        }
                        default /// @src 0:24978:25146  "if (feeIndex == 2) {..."
                        {
                            /// @src 0:25010:25022  "return 10000"
                            var := /** @src 0:25017:25022  "10000" */ 0x2710
                            /// @src 0:25010:25022  "return 10000"
                            leave
                        }
                    }
                    default /// @src 0:24919:25146  "if (feeIndex == 1) {..."
                    {
                        /// @src 0:24951:24962  "return 3000"
                        var := /** @src 0:24958:24962  "3000" */ 0x0bb8
                        /// @src 0:24951:24962  "return 3000"
                        leave
                    }
                }
                default /// @src 0:24861:25146  "if (feeIndex == 0) {..."
                {
                    /// @src 0:24893:24903  "return 500"
                    var := /** @src 0:24900:24903  "500" */ 0x01f4
                    /// @src 0:24893:24903  "return 500"
                    leave
                }
            }
            /// @ast-id 1534 @src 0:25165:25338  "function getV2Factory(uint index) internal pure returns (address){..."
            function fun_getV2Factory(var_index) -> var_
            {
                /// @src 0:25222:25229  "address"
                var_ := /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ 0
                /// @src 0:25241:25331  "if (index ==0){..."
                if /** @src 0:25245:25254  "index ==0" */ iszero(var_index)
                /// @src 0:25241:25331  "if (index ==0){..."
                {
                    /// @src 0:25270:25319  "return 0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"
                    var_ := /** @src 0:25277:25319  "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f" */ 0x5c69bee701ef814a2b6a3edd4b1652cb9cc5aa6f
                    /// @src 0:25270:25319  "return 0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"
                    leave
                }
            }
            /// @ast-id 1549 @src 0:25344:25517  "function getV3Factory(uint index) internal pure returns (address){..."
            function fun_getV3Factory(var_index) -> var
            {
                /// @src 0:25401:25408  "address"
                var := /** @src 0:2137:25520  "contract UniswapMixedExecutor {..." */ 0
                /// @src 0:25420:25510  "if (index ==0){..."
                if /** @src 0:25424:25433  "index ==0" */ iszero(var_index)
                /// @src 0:25420:25510  "if (index ==0){..."
                {
                    /// @src 0:25449:25498  "return 0x1F98431c8aD98523631AE4a59f267346ea31F984"
                    var := /** @src 0:25456:25498  "0x1F98431c8aD98523631AE4a59f267346ea31F984" */ 0x1f98431c8ad98523631ae4a59f267346ea31f984
                    /// @src 0:25449:25498  "return 0x1F98431c8aD98523631AE4a59f267346ea31F984"
                    leave
                }
            }
        }
        data ".metadata" hex"a2646970667358221220df33162013cf21188556ce8d15646f33d90eb20c5ff09e03450ce539fe8b4e8f64736f6c63430008130033"
    }
}
