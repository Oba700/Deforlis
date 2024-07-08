#include <iostream>
#include <string>
#include <vector>

class Binding
{
public:
    std::string bindingName;
    std::string IPv4addr;
    std::string IPv4port;
    Binding(std::string bindingSection)
    {
        while (bindingSection.length() > 0)
        {
            int newLineIndex = bindingSection.find("\n");
            std::string sectionLine = bindingSection.substr(0, newLineIndex);
            int commentStartIndex = std::min(sectionLine.find("#"), sectionLine.find(";"));
            if (commentStartIndex > -1)
            {
                sectionLine = sectionLine.substr(0, commentStartIndex);
            }
            if (sectionLine.length() > 0)
            {
                // std::cout << "Section line: ";
                // std::cout << sectionLine + "\n";
                int equalsSignIndex = sectionLine.find("=");
                if (equalsSignIndex < 1)
                {
                    throw std::invalid_argument("Config line " + sectionLine + " sucks at being meaningful");
                }
                std::string lineKey = sectionLine.substr(0, equalsSignIndex);
                // std::cout << "Key: " + lineKey + "\n";
                std::string lineValue = sectionLine.substr(equalsSignIndex + 1, sectionLine.length());
                // std::cout << "Value: " + lineValue + "\n";
                if (lineKey == "IPv4addr")
                {
                    IPv4addr = lineValue;
                }
                else if (lineKey == "IPv4port")
                {
                    IPv4port = lineValue; // std::stoi( lineValue );
                }
                else if (lineKey == "bindingName")
                {
                    bindingName = lineValue;
                }
                else
                {
                    throw std::invalid_argument("Config line " + sectionLine + " has key that sucks at being valid key");
                }
            }
            bindingSection = bindingSection.substr(newLineIndex + 1, bindingSection.length());
        }
    }
    Binding() = default;
};

class Config
{
public:
    std::vector<Binding> bindings;
    Config(std::string configFileContent)
    {
        while (true)
        {
            const std::string bindingSectionHeader = "[binding]";
            int bindingSectionStart = configFileContent.find(bindingSectionHeader);
            if (bindingSectionStart == -1)
            {
                if (this->bindings.size() == 0)
                {
                    throw std::invalid_argument("Config sucks at having any [binding] header. Must have one or more");
                }
                else
                {
                    break;
                }
            }

            std::string noBindingHeaderConf = configFileContent.substr(bindingSectionStart + bindingSectionHeader.length());
            int bindingSectionFinish = noBindingHeaderConf.find("[");
            if (bindingSectionFinish == -1)
            {
                bindingSectionFinish = configFileContent.length();
            }
            std::string bindingSection = configFileContent.substr(bindingSectionStart + bindingSectionHeader.length(), bindingSectionFinish);
            bindings.push_back(bindingSection);
            configFileContent = configFileContent.substr(bindingSectionFinish, configFileContent.length());
        }
    }
};
