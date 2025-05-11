# KNIGHTS BOARD
Sample golang exercise

## USAGE
1. Clone the repo
2. Run build command
```
docker build -t <IMAGE_NAME>:latest .
```
3. Run built image
```
docker run -e BOARD_API=<BOARD_DATA_ENDPOINT> -e COMMANDS_API=<COMMANDS_DATA_ENDPOINT> <IMAGE_NAME>:latest
```