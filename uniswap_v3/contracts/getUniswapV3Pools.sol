// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

//import "forge-std/Test.sol";

interface IUniswapV2Facotry {
    function getPair(address tokenA, address tokenB) external view returns (address pair);
    function allPairs(uint) external view returns (address pair);
    function allPairsLength() external view returns (uint);
}
struct PopulatedTick {
    int24 tick;
    int128 liquidityNet;
    uint128 liquidityGross;
}
interface TickLens{

    function getPopulatedTicksInWord(address pool, int16 tickBitmapIndex) external view returns (PopulatedTick[] memory populatedTicks);
}

interface IUniswapV3Pool{
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


//ticklens address

contract GetUniswapV3Pools {
    constructor(
        address[] memory pools 
    ) {



        uint len = pools.length;

        //address[] memory token0Arr = new address[](len);
        //address[] memory token1Arr = new address[](len);
        uint128[] memory liquidityArr = new uint128[](len);
        uint160[] memory sqrtPriceX96Arr = new uint160[](len);
        int24[] memory tickCurrentArr = new int24[](len);
        int24[] memory tickSpacingArr = new int24[](len);
        

        for (uint256 i = 0; i < len; ++i) {
            IUniswapV3Pool poolInstance = IUniswapV3Pool(pools[i]);
            //token0Arr[i] = poolInstance.token0();
            //token1Arr[i] = poolInstance.token1();
            (uint160 sqrtPriceX96, int24 tick, , , , , ) = poolInstance.slot0();
            liquidityArr[i] = poolInstance.liquidity();
            sqrtPriceX96Arr[i] = sqrtPriceX96;
            tickCurrentArr[i] = tick;
            tickSpacingArr[i] = poolInstance.tickSpacing();


            //console.log("liquidity", poolInstance.liquidity());
            //console.log("sqrtPriceX96", sqrtPriceX96);
        }  


        bytes memory data = abi.encode(block.number, liquidityArr, sqrtPriceX96Arr, tickCurrentArr, tickSpacingArr);
        //bytes memory data = abi.encode(block.number, liquidityArr, sqrtPriceX96Arr,tickCurrentArr, populatedTicksArr);

        assembly {
            return(add(data, 32), mload(data))
        }
    }
}
