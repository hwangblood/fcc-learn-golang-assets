module github.com/hwangblood/hellogo

go 1.20

// ! replace keyword is not suitable for production, we shoul push it to remote repository
replace "github.com/hwangblood/mystrings" v0.0.0 => "../mystrings"

require ( 
    "github.com/hwangblood/mystrings" v0.0.0
)