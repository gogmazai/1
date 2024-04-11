import Data.Char

rta :: Char -> Int
rta 'I' = 1
rta 'V' = 5
rta 'X' = 10
rta 'L' = 50
rta 'C' = 100
rta 'D' = 500
rta 'M' = 1000
rta _   = 0

rom_to_ar :: String -> Int
rom_to_ar s = rom_to_ar' s 0
  where
    rom_to_ar' [] acc = acc
    rom_to_ar' [x] acc = acc + rta x
    rom_to_ar' (x:y:xs) acc
      | rta x < rta y = rom_to_ar' (y:xs) (acc - rta x)
      | otherwise = rom_to_ar' (y:xs) (acc + rta x)

at_to_rom :: Int -> String
at_to_rom n
    | n >= 100 = "C" ++ at_to_rom (n - 100)
    | n >= 90 = "XC" ++ at_to_rom (n - 90)
    | n >= 50 = "L" ++ at_to_rom (n - 50)
    | n >= 40 = "XL" ++ at_to_rom (n - 40)
    | n >= 10 = "X" ++ at_to_rom (n - 10)
    | n >= 9 = "IX" ++ at_to_rom (n - 9)
    | n >= 5 = "V" ++ at_to_rom (n - 5)
    | n >= 4 = "IV" ++ at_to_rom (n - 4)
    | n >= 1 = "I" ++ at_to_rom (n - 1)
    | otherwise = ""


romans = ["I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"]
arabic = map show [1..10]
operations = ["+", "-", "*", "/"]


g op a b
  |op == "+" = a + b
  |op == "-" = a - b
  |op == "*" = a * b
  |op == "/" = div a b

result op a b list
  |length list == 3 && elem op operations && elem a arabic && elem b arabic  =   putStrLn $ show (g op (read a) (read b))
  |length list == 3 && op == "-" && elem a romans && elem b romans && rom_to_ar a > rom_to_ar b = putStrLn (at_to_rom (g op (rom_to_ar a) (rom_to_ar b))) 
  |length list == 3 && op == "/" && elem a romans && elem b romans && rom_to_ar a >= rom_to_ar b = putStrLn (at_to_rom (g op (rom_to_ar a) (rom_to_ar b)))
  |length list == 3 && op == "*" || op == "+" && elem a romans && elem b romans = putStrLn (at_to_rom (g op (rom_to_ar a) (rom_to_ar b))) 
  |otherwise = putStrLn "Неверный формат ввода"  
 

main :: IO()
main = do
  input <- getLine
  let primer = words input :: [String]
  let a = (words input) !! 0 :: String
  let b = (words input) !! 2 :: String
  let op = (words input) !! 1 :: String
  result op a b primer


