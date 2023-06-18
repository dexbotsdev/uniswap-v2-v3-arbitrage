// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

interface IUniswapV2Facotry {
    function getPair(address tokenA, address tokenB) external view returns (address pair);
    function allPairs(uint) external view returns (address pair);
    function allPairsLength() external view returns (uint);
}

interface IUniswapV3Pool{
    mapping(int24 => Tick.Info) public override ticks;
    function token0() external view returns (address);
    function token1() external view returns (address);
    function slot0() external view returns (uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, uint8 feeProtocol, bool unlocked);
    function fee() external view returns (uint24);
    function tickSpacing() external view returns (int24);
    function observationCardinalityNext() external view returns (uint16);
    function observationCardinality() external view returns (uint16);
    function observationIndex() external view returns (uint16);
    function observation(uint256 index) external view returns (uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulativeX128, bool initialized);
    function liquidity() external view returns (uint128);
    function sqrtPriceX96() external view returns (uint160);
    function tickCumulative(uint256 index) external view returns (int56);
    function secondsPerLiquidityCumulativeX128(uint256 index) external view returns (uint160);
    function tick() external view returns (int24);
    function maxLiquidityPerTick() external view returns (uint128);
    function liquidityCumulative(uint256 index) external view returns (uint256);
    function initialize(uint160 sqrtPriceX96) external;
    function mint(address recipient, int24 tickLower, int24 tickUpper, uint128 amount, bytes calldata data) external returns (uint256 amount0, uint256 amount1);
    function burn(int24 tickLower, int24 tickUpper, uint128 amount) external returns (uint256 amount0, uint256 amount1);
    function swap(address recipient, bool zeroForOne, int256 amountSpecified, uint160 sqrtPriceLimitX96, bytes calldata data) external returns (int256 amount0, int256 amount1);
    function flash(address recipient, uint256 amount0, uint256 amount1, bytes calldata data) external;
    function increaseObservationCardinalityNext(uint16 observationCardinalityNext) external;
    function decreaseObservationCardinalityNext(uint16 observationCardinalityNext) external;
    function initialize(uint160 sqrtPriceX96, uint16 _observationCardinalityLimit, uint16 _observationCardinalityNext) external;
    function collectProtocol(address token0, address token1, uint128 amount0Requested, uint128 amount1Requested) external returns (uint128 amount0, uint128 amount1);
}

contract GetUniswapV3Pools {
    constructor(
        address factory,
        uint start,
        uint end
    ) {

        //  Token0       common.Address
        //Token1       common.Address
        //Fee          constants.FeeAmount
        //SqrtRatioX96 *big.Int
        //Liquidity    *big.Int
        //TickCurrent  int

        //TickInfoMap map[int]*Tick

        address[] memory pools = new address[](end-start);
        address[] memory token0s = new address[](end-start);
        address[] memory token1s = new address[](end-start);
        uint112[] memory reserve0s = new uint112[](end-start);
        uint112[] memory reserve1s = new uint112[](end-start);
        mapping(int24 => Tick.Info) public override ticks;

        for (uint256 i = start; i < end; ++i) {
            address pair = IUniswapV2Facotry(factory).allPairs(i);
            pairs[i-start] = pair;
            token0s[i-start] = IUniswapV2Pair(pair).token0();
            token1s[i-start] = IUniswapV2Pair(pair).token1();
            (uint112 reserve0, uint112 reserve1,) = IUniswapV2Pair(pair).getReserves();
            reserve0s[i-start] = reserve0;
            reserve1s[i-start] = reserve1;

        }  


        bytes memory data = abi.encode(block.number, pairs, token0s, token1s, reserve0s, reserve1s);
        assembly {
            return(add(data, 32), mload(data))
        }
    }
}
