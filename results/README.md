Here you can find the raw benchmark and analysis summary data.

The naming scheme consists of 4 test variables:

1. The system it was tested on 
    - `T14s` being my ThinkPad T14s with an AMD Ryzen 7 PRO 5850U & 32GiB RAM
    - `PC` being my workstation PC with an AMD Ryzen 7 5800X & 32GiB RAM
2. The number of items per slice (`s10k` means `len(s) == 10_000`)
3. The number of test runs (`n50` means 50 runs per benchmark)
4. The methodology
    - Where no suffix is present, the tests were performed before Go v1.22 was released and thus tested with the `GOEXPERIMENT=loopvar` environment variable specified.
    - With the suffix `swap`, the two tests were actually performed using the latest release of Go v1.21 and Go v1.22.

Only the latest two tets also contain the raw output of the tests, because I am stupid and only saved them in these tests. 🤦