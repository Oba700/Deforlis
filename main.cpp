#include <iostream>
#include <fstream>
#include <thread>
#include <string>
#include "config.cpp"
#include "listen.cpp"

const std::string CONFIG_FILE_CLI_ARG_NAME = "--configfile=";
const std::string PRODUCT_NAME = "deforlis";

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
        configFileCliArgument = PRODUCT_NAME + ".ini";
        std::cout << "Nothing specified with " + CONFIG_FILE_CLI_ARG_NAME + ". Using " + configFileCliArgument + " instead.\n";
    }
    std::string configFileContent = configFile(configFileCliArgument);
    std::cout << "Hello from " + PRODUCT_NAME + " web server\n";
    Config config = Config(configFileContent);
    //std::cout << config.bindings.size() << "\n";
    for (auto i = config.bindings.begin(); i != config.bindings.end(); ++i)
    {
        std::cout << (*i).bindingName << " " << (*i).IPv4addr << ":" << (*i).IPv4port <<  "\n";
        listen((*i).IPv4addr,(*i).IPv4port);
    }
    return 0;
}
