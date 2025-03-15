#include <iostream>
using namespace std;

int main() {
    // Array of strings to store ASCII art lines
    string asciiArt[] = {
        "  _   _      _ _         __        __         _     _ ",
        " | | | |    | | |        \\ \\      / /        | |   | |",
        " | |_| | ___| | | ___     \\ \\ /\\ / /__  _ __ | | __| |",
        " |  _  |/ _ \\ | |/ _ \\     \\ V  V / _ \\| '_ \\| |/ _` |",
        " | | | |  __/ | | (_) |     \\_/\\_/ / (_) | | | | (_| |",
        " \\_| |_/\\___|_|_|\\___( )     \\_/\\_/ \\___/|_| |_|\\__,_|",
        "                    |/                                "
    };

    // Loop through the array and print each line
    for (int i = 0; i < 7; ++i) {
        cout << asciiArt[i] << endl;
    }

    return 0;
}
