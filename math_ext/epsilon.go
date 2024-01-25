package math_ext

// https://en.wikipedia.org/wiki/Machine_epsilon

// EPSILON_FLOAT_32 represents upper bound on the relative approximation error due to rounding in 32-bit floating point arithmetic.
const EPSILON_FLOAT_32 = float32(7.)/3 - float32(4.)/3 - float32(1.)

// EPSILON_FLOAT_64 represents upper bound on the relative approximation error due to rounding in 64-bit floating point arithmetic.
const EPSILON_FLOAT_64 = float64(7.)/3 - float64(4.)/3 - float64(1.)
