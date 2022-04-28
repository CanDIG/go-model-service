# Debugging a containerized process with VSCode

This is a short guide for using VSCode to code and debug right inside of a running Docker container. This allows you to get access to the filesystem of your container, which is very helpful if your build process uses code-generation, or if you want to debug someone else's dockerized software without downloading and installing all of their code into your local working directories.

## Attaching to containers
1) Install `Remote - Containers` extension in VSCode. Reload VSCode if needed.
2) Spin up the container with `docker run` or `docker-compose up`.
3) Find your container in the `Remote Explorer` browser on the left side of the screen (it's usually right above the `Extensions` browser tab).
4) Right-click the container you want to attach to and select `Attach in New Window`
5) Open the folder for your working directory (called an `Active Folder` in VSCode). The path will be the absolute path *in the container*.

## Extensions
You must install whatever extensions you need for development into your container, eg. `Docker`, `Go`, etc. Even if you have these extensions installed locally, they will *not* be installed in your container by default. Note that in the `Extensions` browser, the top panel shows your `Local - Installed` extensions, and you can hit the cloud-shaped download button to quickly download all of them into your container.
After the initial connection, your container will appear under `Dev Containers` in the `Remote Explorer`. Next time you connect to a `Dev Container`, VSCode will install that container's preferred extensions into it automatically, even if the container was rebuilt.

## Debugging containers
1) If needed, launch your remote debugger inside your container (use the VSCode `Terminal` or `docker exec`) and attach it to your application.
2) In VSCode, select `Run > Add Configuration...`. Modify the generated `.vscode/launch.json` to `attach` to the port that your debugger is running at (usually `127.0.0.1:2345`). You may also want to push this file to source control to speed up future set-up. My `.vscode/launch.json` looks like:
```
{
    "version": "0.2.0",
    "configurations": [
    
        {
            "name": "Attach",
            "type": "go",
            "request": "attach",
            "mode": "remote",
            "remotePath": "${workspaceFolder}",
            "port": 2345,
            "host": "127.0.0.1",
            "apiVersion": 1
        }
    ]
}
```
3) When you're ready to debug, hit `Run > Start Debugging`.

## Tips and Tricks
### Prerequisites for debugging
If you are unable to set breakpoints, check that you have the extension for your language installed in the container.

Similarly, any tool prerequisites for debugging (eg. `delve` for Go) must be installed in your container. You can use the Terminal in VSCode to install tools into your container as you go, or alternately run `docker exec`/`docker-compose exec`.

In order to run debuggers, it may be necessary to run the docker container in `priviledged` mode. If you choose to do this, first understand the risks of giving your container root permissions, and make sure that you are only using the `--priviledged` flag/`priviledged: true` configuration where appropriate (*not* in production containers.)

#### Difficulty attaching to a process/pid from within the container
If you are seeing errors relating to `.../yama/ptrace_scope`, see [this article](https://bitworks.software/en/2017-07-24-docker-ptrace-attach.html) for help. Be risk-aware and careful if you choose to disable certain security measures.

### Containerize the development environment
Consider maintaining a seperate `development.Dockerfile` for building containers for development and debugging. The development container can include files for source control (`.git`), VSCode configuration (`.vscode`), and development scripts that are not needed in testing or production containers.

You could even set up source control to allow you to push code *from within your development container*, eliminating the need to map your container volumes to your local directory for active development. This is very helpful if your build process, like mine, contains a lot of code-generation that you prefer to keep away from source control.

## Works Cited
* [How to debug a running Go app with VSCode](https://medium.com/average-coder/how-to-debug-a-running-go-app-with-vscode-76e3eac45bd)
* [The solution for enabling of ptrace and PTRACE_ATTACH in Docker Containers](https://bitworks.software/en/2017-07-24-docker-ptrace-attach.html)

