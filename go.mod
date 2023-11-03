module github.com/derrynEdwards/qualys-toolbox

go 1.21.1

require github.com/joho/godotenv v1.5.1

replace github.com/derrynEdwards/qualys-toolbox/internal/qualyscache v0.0.0 => ./internal/qualyscache

replace github.com/derrynEdwards/qualys-toolbox/internal/qualysapi v0.0.0 => ./internal/qualysapi

replace github.com/derrynEdwards/qualys-toolbox/cmd/qualysToolbox v0.0.0 => ./cmd/qualysToolbox

replace github.com/derrynEdwards/qualys-toolbox/pkg/config v0.0.0 => ./pkg/config
