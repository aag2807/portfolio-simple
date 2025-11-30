(function() {
    var canvas = document.getElementById('particle-canvas');
    if (!canvas) return;

    var ctx = canvas.getContext('2d');
    if (!ctx) return;

    var gridWidth = 200;
    var gridHeight = 150;
    var cellSize;
    var grid = [];
    var nextGrid = [];
    var mouse = { x: -1, y: -1, down: false };
    var brushSize = 3;

    function isDark() {
        return document.documentElement.classList.contains('dark');
    }

    function getColors() {
        if (isDark()) {
            return {
                sand: [
                    [92, 130, 184],
                    [108, 150, 200],
                    [130, 170, 220],
                    [150, 185, 230]
                ],
                bg: 'transparent'
            };
        }
        return {
            sand: [
                [71, 85, 105],
                [85, 100, 120],
                [100, 115, 135],
                [115, 130, 150]
            ],
            bg: 'transparent'
        };
    }

    function resize() {
        var section = canvas.parentElement;
        canvas.width = section.offsetWidth;
        canvas.height = section.offsetHeight;
        cellSize = Math.max(canvas.width / gridWidth, canvas.height / gridHeight);
        gridWidth = Math.floor(canvas.width / cellSize);
        gridHeight = Math.floor(canvas.height / cellSize);
        initGrid();
    }

    function initGrid() {
        grid = [];
        nextGrid = [];
        for (var y = 0; y < gridHeight; y++) {
            grid[y] = [];
            nextGrid[y] = [];
            for (var x = 0; x < gridWidth; x++) {
                grid[y][x] = 0;
                nextGrid[y][x] = 0;
            }
        }
    }

    function getCell(x, y) {
        if (x < 0 || x >= gridWidth || y < 0 || y >= gridHeight) return 1;
        return grid[y][x];
    }

    function setNextCell(x, y, val) {
        if (x < 0 || x >= gridWidth || y < 0 || y >= gridHeight) return;
        nextGrid[y][x] = val;
    }

    function spawnSand(cx, cy, radius) {
        for (var dy = -radius; dy <= radius; dy++) {
            for (var dx = -radius; dx <= radius; dx++) {
                if (dx * dx + dy * dy <= radius * radius) {
                    var x = cx + dx;
                    var y = cy + dy;
                    if (x >= 0 && x < gridWidth && y >= 0 && y < gridHeight) {
                        if (Math.random() > 0.5) {
                            var colorIndex = Math.floor(Math.random() * 4) + 1;
                            grid[y][x] = colorIndex;
                        }
                    }
                }
            }
        }
    }

    function update() {
        for (var y = 0; y < gridHeight; y++) {
            for (var x = 0; x < gridWidth; x++) {
                nextGrid[y][x] = 0;
            }
        }

        for (var y = gridHeight - 1; y >= 0; y--) {
            for (var x = 0; x < gridWidth; x++) {
                var cell = grid[y][x];
                if (cell === 0) continue;

                var below = getCell(x, y + 1);
                var belowLeft = getCell(x - 1, y + 1);
                var belowRight = getCell(x + 1, y + 1);

                if (y >= gridHeight - 3) {
                    if (Math.random() < 0.02) {
                        continue;
                    }
                }

                if (y === gridHeight - 1) {
                    setNextCell(x, y, cell);
                } else if (below === 0) {
                    setNextCell(x, y + 1, cell);
                } else if (belowLeft === 0 && belowRight === 0) {
                    if (Math.random() > 0.5) {
                        setNextCell(x - 1, y + 1, cell);
                    } else {
                        setNextCell(x + 1, y + 1, cell);
                    }
                } else if (belowLeft === 0) {
                    setNextCell(x - 1, y + 1, cell);
                } else if (belowRight === 0) {
                    setNextCell(x + 1, y + 1, cell);
                } else {
                    if (y > gridHeight - 15 && Math.random() < 0.005) {
                        continue;
                    }
                    setNextCell(x, y, cell);
                }
            }
        }

        var temp = grid;
        grid = nextGrid;
        nextGrid = temp;
    }

    function draw() {
        var colors = getColors();
        ctx.clearRect(0, 0, canvas.width, canvas.height);

        for (var y = 0; y < gridHeight; y++) {
            for (var x = 0; x < gridWidth; x++) {
                var cell = grid[y][x];
                if (cell > 0) {
                    var color = colors.sand[cell - 1];
                    ctx.fillStyle = 'rgba(' + color[0] + ',' + color[1] + ',' + color[2] + ',0.8)';
                    ctx.fillRect(x * cellSize, y * cellSize, cellSize, cellSize);
                }
            }
        }
    }

    var frameCount = 0;
    var autoSpawnX = gridWidth * 0.3;

    function animate() {
        frameCount++;

        if (frameCount % 2 === 0) {
            autoSpawnX = gridWidth * (0.25 + 0.25 * Math.sin(frameCount * 0.02));
            spawnSand(Math.floor(autoSpawnX), 2, 2);
            
            var secondX = gridWidth * (0.75 - 0.2 * Math.sin(frameCount * 0.015));
            spawnSand(Math.floor(secondX), 2, 1);
        }

        if (mouse.down) {
            var gridX = Math.floor(mouse.x / cellSize);
            var gridY = Math.floor(mouse.y / cellSize);
            spawnSand(gridX, gridY, brushSize);
        }

        update();
        draw();
        requestAnimationFrame(animate);
    }

    canvas.addEventListener('mousemove', function(e) {
        var rect = canvas.getBoundingClientRect();
        mouse.x = e.clientX - rect.left;
        mouse.y = e.clientY - rect.top;
    });

    canvas.addEventListener('mousedown', function(e) {
        mouse.down = true;
        var rect = canvas.getBoundingClientRect();
        mouse.x = e.clientX - rect.left;
        mouse.y = e.clientY - rect.top;
    });

    canvas.addEventListener('mouseup', function() {
        mouse.down = false;
    });

    canvas.addEventListener('mouseleave', function() {
        mouse.down = false;
    });

    canvas.addEventListener('touchstart', function(e) {
        e.preventDefault();
        mouse.down = true;
        var rect = canvas.getBoundingClientRect();
        var touch = e.touches[0];
        mouse.x = touch.clientX - rect.left;
        mouse.y = touch.clientY - rect.top;
    }, { passive: false });

    canvas.addEventListener('touchmove', function(e) {
        e.preventDefault();
        var rect = canvas.getBoundingClientRect();
        var touch = e.touches[0];
        mouse.x = touch.clientX - rect.left;
        mouse.y = touch.clientY - rect.top;
    }, { passive: false });

    canvas.addEventListener('touchend', function() {
        mouse.down = false;
    });

    window.addEventListener('resize', resize);

    resize();
    animate();
})();
