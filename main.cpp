#include <iostream>
#include <fstream>
#include <string>
#include "config.cpp"

const std::string CONFIG_FILE_CLI_ARG_NAME = "--configfile=";
const std::string PRODUCT_NAME = "shity";


std::string configFile(std::string filePath)
{
    std::string line;
    std::string content;
    std::ifstream f(filePath);
    if (f.good())
    {
        while (getline(f, line))
        {
            content += line;
            content += "\n";
        }
        f.close();
        return content;
    }
    else
    {
        throw std::invalid_argument(filePath + " sucks at being a file");
    }
}

int main(int argc, char *argv[])
{
    std::string configFileCliArgument;
    for (int i = 0; i < argc; i++)
    {
        std::string argOption = argv[i];
        if (argOption.rfind(CONFIG_FILE_CLI_ARG_NAME, 0) == 0)
        {
            configFileCliArgument = argOption.substr(CONFIG_FILE_CLI_ARG_NAME.length(), argOption.length());
            break;
        }
    }
    if (configFileCliArgument.length() == 0)
    {
        configFileCliArgument = PRODUCT_NAME + ".json";
        std::cout << "Nothing specified with " + CONFIG_FILE_CLI_ARG_NAME + ". Using " + configFileCliArgument + " instead.\n";
    }
    std::string configFileContent = configFile(configFileCliArgument);
    std::cout << "Hello from " + PRODUCT_NAME + " web server\n";
    Config config = Config(configFileContent);
    std::cout << "Gonna bind to " + config.binding.IPv4addr + ":" + config.binding.IPv4port + "\n";
    return 0;
}
