pragma solidity 0.8.19;

library Utils {
    bytes32 private constant POOL_INIT_CODE_HASH = 0x96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f;

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

