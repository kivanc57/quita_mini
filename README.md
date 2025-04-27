
# QUITA Mini

QUITA Mini is a text analysis tool that calculates various linguistic metrics from text data. It processes a collection of text files, calculates metrics like Type-Token Ratio (TTR), entropy, average token lengths, hapax legomena percentages, and other related linguistic statistics, and outputs the results in an Excel file.

## Features

- **Text Analysis**: Tokenizes and processes text to calculate various linguistic metrics.
- **Metrics Calculated**:
  - Type-Token Ratio (TTR)
  - Entropy
  - Normalized Entropy
  - Average Token Length
  - Average Type Length
  - Token Length Standard Deviation
  - Type Length Standard Deviation
  - Yule's K
  - Adjusted Modulus
  - Percentage of Text Hapax
  - Percentage of Type Hapax
  - R1, R Index, Rr Index, and L Index
- **Excel Output**: Saves results in an Excel file for further analysis or visualization.

## Installation

1. Clone this repository to your local machine.
2. Ensure that you have Go installed. If not, you can download and install it from [here](https://golang.org/dl/).
3. Install the necessary Go packages:
    ```bash
    go get github.com/xuri/excelize/v2
    ```
4. Install other dependencies by running:
    ```bash
    go mod tidy
    ```

## Usage

1. Set the `absDirPath` variable in `main.go` to the path of the directory containing your text files.
2. Run the program using:
    ```bash
    go run main.go
    ```

3. After execution, the results will be saved to `results.xlsx` in the current directory.


## License

This project is licensed under the GNU 3.0 License - see the [GNU GENERAL PUBLIC LICENSE](LICENSE) file for details.

## Acknowledgments

- This project uses the [Excelize library](https://github.com/xuri/excelize) for handling Excel files.
