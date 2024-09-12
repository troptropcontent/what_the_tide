git config --global user.email $GIT_EMAIL
git config --global user.name $GIT_USER
git config --global init.defaultBranch main
git config --global pull.rebase false
git config --global push.autoSetupRemote true

# Install KAMAL for deployments
gem install kamal

# download SASS if it is not there already
bin_directory="$(pwd)/bin"
sass_path="$bin_directory/sass"

if [ ! -f $sass_path ]; then
    echo "Sass binary not found, downloading it."
    # Download the binary
    curl https://github.com/sass/dart-sass/releases/download/1.78.0/dart-sass-1.78.0-linux-x64.tar.gz -o $sass_path
    # Grant execution permission to the user on the binary
    chmod u+x $sass_path
    echo "Sass binary downloaded."
else
    echo "Sass binary found, no need to download it."
fi

