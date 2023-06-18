from symtable import Symbol
from sympy import *


a,b,c,d,A,B,C,D,r = symbols('a b c d A B C D r')

Ea, Eb, R0, R1, R1p, R2, deltaA, deltaAPrime = symbols('Ea Eb R0 R1 R1p R2 deltaA deltaAPrime')

#R0 -> A
#R1 -> B
#R1' -> C
#R2 -> D






Ea = (R0*R1p)/(R1p+R1*r)

Eb = (r*R1*R2)/(R1p+R1*r)

deltaAPrime = (Ea*r*deltaA)/(Eb+r*deltaA)

f = deltaAPrime-deltaA







# b = (B*r*a)/(A+r*a)
# f= b-a
# der= diff(f,a)
# print(f)
# print("der",simplify(der))
# print(solve(der,a))



# reserve0 1.643186143e+22
# reserve1 3.008006013e+13
# reserve2 6753022
# reserve3 6.75346784e+18
# virtualReserve0 3.70007879e+15
# virtualReserve1 6.753466319e+18

fee = 0.997


reserve0 = 3.70007879e+15
reserve1 = 6.753466319e+18
reserve2 = 1.274451364e+20
reserve3 = 7.518656714e+16

virtualReserve0 =reserve0*reserve2/(reserve2+reserve1*0.997)
virtualReserve1 =reserve1*reserve3*fee/(reserve2+reserve1*0.997)

print("virtual reserve 0: ",reserve0*reserve2/(reserve2+reserve1*0.997))
print("virtual reserve 1: ",reserve1*reserve3*fee/(reserve2+reserve1*0.997))

bestAmountIn = ((virtualReserve0*virtualReserve1*fee)**0.5-virtualReserve0)/fee

print("best amount in: ", bestAmountIn)
print("revenue: ", (virtualReserve1*bestAmountIn*fee)/(virtualReserve0 + bestAmountIn*fee)-bestAmountIn)

print(solve(D*r*(B*r*a/(A+r*a))/(C+r*(B*r*a/(A+r*a))),a))
print(solve((r*B*D/(C+B*r))*r*a/((A*C/(C+B*r))+r*a),a))

print(simplify(D*r*(B*r*a/(A+r*a))/(C+r*(B*r*a/(A+r*a)))-(r*B*D/(C+B*r))*r*a/((A*C/(C+B*r))+r*a)))


# round  1
# pair address:  0x3139Ffc91B99aa94DA8A2dc13f1fC36F9BDc98eE
# tokenAddress0:  0x8E870D67F660D95d5be530380D0eC0bd388289E1
# tokenAddress1:  0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48
# edge reserve0:  BigNumber { s: 1, e: 18, c: [ 89853, 23276609552323 ] }
# edge reserve1:  BigNumber { s: 1, e: 6, c: [ 4980878 ] }
# reserve 0:  342294163068755572
# reserve 1:  1.145354321615050090057e+21
# reserve 2:  8985323276609552323
# reserve 3:  4980878
# virturalReserve0:  2672355668335084.68352692569491121545
# virturalReserve1:  4941991.33761215248881509883

# round  2
# pair address:  0xAE461cA67B15dc8dc81CE7615e0320dA1A9aB8D5
# tokenAddress0:  0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48
# tokenAddress1:  0x6B175474E89094C44Da98b954EedeAC495271d0F
# edge reserve0:  BigNumber { s: 1, e: 13, c: [ 47249143980502 ] }
# edge reserve1:  BigNumber { s: 1, e: 25, c: [ 472336399231, 83288662457743 ] }
# reserve 0:  2672355668335084.68352692569491121545
# reserve 1:  4941991.33761215248881509883
# reserve 2:  47249143980502
# reserve 3:  4.7233639923183288662457743e+25
# virturalReserve0:  2672355389660471.64510601268286588465
# virturalReserve1:  4925548078781681296.6354668797682183268

# round  3
# pair address:  0x095739e9Ea7B0d11CeE1c1134FB76549B610f4F3
# tokenAddress0:  0x6B175474E89094C44Da98b954EedeAC495271d0F
# tokenAddress1:  0xB6eD7644C69416d67B522e20bC294A9a9B405B31
# edge reserve0:  BigNumber { s: 1, e: 7, c: [ 32047005 ] }
# edge reserve1:  BigNumber { s: 1, e: 0, c: [ 1 ] }
# reserve 0:  2672355389660471.64510601268286588465
# reserve 1:  4925548078781681296.6354668797682183268
# reserve 2:  32047005
# reserve 3:  1
# virturalReserve0:  17439.41612334402453042244
# virturalReserve1:  0.99999999999347414038

# round  4
# pair address:  0xc12c4c3E0008B838F75189BFb39283467cf6e5b3
# tokenAddress0:  0xB6eD7644C69416d67B522e20bC294A9a9B405B31
# tokenAddress1:  0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2
# edge reserve0:  BigNumber { s: 1, e: 12, c: [ 2776016295639 ] }
# edge reserve1:  BigNumber { s: 1, e: 18, c: [ 85974, 63208228714843 ] }
# reserve 0:  17439.41612334402453042244
# reserve 1:  0.99999999999347414038
# reserve 2:  2776016295639
# reserve 3:  8597463208228714843
# virturalReserve0:  17439.41612333776120340157
# virturalReserve1:  3087759.54666070659049058193

# final virturalReserve0:  17439.41612333776120340157
# final virturalReserve1:  3087759.54666070659049058193
# best amount in:  214910.25081337044055505664
# best revenue:  2640447.15323523174628778905