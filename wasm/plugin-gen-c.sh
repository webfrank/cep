read -p "Plugin name: " plugin
extism gen plugin -l C -o $plugin
cd ./$plugin
git submodule update --init --recursive
