def main():
    f = open("elfs.txt", "r")
    txt = f.read()
    f.close()
    elfs = txt.split("\n\n")
    def a():    
        m = 0
        for elf in elfs:
            elfpacking = [int(e) for e in elf.split("\n")]
            m = max(m, sum(elfpacking))
        return m
    def b():
        topthree = [0,-1,-2]
        for elf in elfs:
            elfsum = sum([int(e) for e in elf.split("\n")])
            if (elfsum > topthree[0]) or (elfsum > topthree[1]) or (elfsum > topthree[2]):
                topthree.pop()
                topthree.append(elfsum)
                topthree.sort(reverse=True)
        return sum(topthree)

    
        









if __name__ == "__main__":
    main()
    