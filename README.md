# Conway's Game of Life

## Usage

```bash
make build && make run
```

### Live rebuild with [Air](https://github.com/cosmtrek/air)

Only use it if you want to make changes to the program and see them live

```bash
go install github.com/cosmtrek/air@latest
```

```bash
air
```

## Roadmap

-   [x] Render the application using [raylib](https://www.raylib.com/)
-   [x] Add a randomizer for the grid
-   [ ] Implement Game of Life logic
-   [ ] Add window controls
    -   [ ] Play, pause
    -   [ ] Step through generations
    -   [ ] Trails switch
    -   [ ] Speed controller
