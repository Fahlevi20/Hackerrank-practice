package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
    // Write your code here
    
    length := float64(len(arr));
    positive := make([]float64,0);
    negative := make([]float64,0);
    zero := make([]float64,0);
    
    for i := 0; i < len(arr); i++ {
        switch {
        case arr[i] > 0:
            positive = append(positive, 1)
        case arr[i] < 0:
            negative = append(negative, 1)
        case arr[i] == 0:
            zero = append(zero, 1)
        }
    }

    // Tampilkan hasil
    fmt.Println(float64(len(positive))/float64(length))
    fmt.Println(float64(len(negative))/float64(length))
    fmt.Println(float64(len(zero))/float64(length))
}  

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    plusMinus(arr)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
