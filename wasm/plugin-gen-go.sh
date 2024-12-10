read -p "Plugin name: " plugin
extism gen plugin -l Go -o $plugin
cd ./$plugin
git submodule update --init --recursive
