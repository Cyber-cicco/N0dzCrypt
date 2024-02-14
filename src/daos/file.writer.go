package daos

import "os"

func WriteToFile(content []byte, path string) {
    f, err := os.OpenFile(path, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
    defer f.Close()
    if err != nil {
        panic(err)
    }
    f.Write(content)
}
