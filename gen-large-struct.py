SIZE = 1000

with open("largestruct.go", "w") as f:
    f.write("package main\n\n")
    f.write("type LargeStruct struct {\n")
    for i in range(0, SIZE):
        f.write(f"\tF{i} int64\n")
    f.write("}\n")
