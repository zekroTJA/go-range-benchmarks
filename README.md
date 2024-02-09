# Go v1.22 range change benchmarks

In the new [version 1.22]((https://tip.golang.org/doc/go1.22)) of Go, a critical issue with loops, which can cause nasty bugs and unexpected behavior when capturing the reference of a variable in a range loop or when using the value in a closure, has finally been fixed.

**But does this have a performance impact?**

To test this, I've created this little suite of benchmarks. In [`src/`](src/), you can find the benchmark code. In [`results/`](results/), you can find the raw benchmark data and analysis summaries.

The [notebook](notebook.ipynb) contains some graphs and stats of the analyzed data.

If you want to go more in depth, I've written a [blog post](https://blog.zekro.de/how-the-new-loops-in-go-might-be-a-performance-trap) about this topic.

