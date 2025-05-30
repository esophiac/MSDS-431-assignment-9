## import packages to use
library(boot)

## create functions to use in boot to calculate mean in the boots function
samplemean <- function(x, d) {
  return(mean(x[d], trim = 0.05))
}

## set up the data to be used to demonstrate that the boot method works
## population mean: 10, standard deviation: 5, size: 100
set.seed(123)
data_norm <- round(rnorm(100, 10, 5))

## base mean with trimming of 0.05
base_mean <- mean(data_norm, trim = 0.05)

cat("Base Trimmed Mean: ", base_mean)

## calculate the boot sampling mean
boot_mean <- boot(data_norm, samplemean, R=100)

## calculate the mean
calc_mean <- boot_mean$t0

cat("Bootstrap Trimmed Mean: ", calc_mean)
