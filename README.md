# Chess Puzzles in Terminal

# Set up
## Step 1
```
git clone https://github.com/MilesJuddPorter/chessTerm.git
cd chessTerm
go build -o chessTerm
sudo mv chessTerm /usr/local/bin/
```
## Step 2
- ```cd puzzles; pwd```
- Copy this so you can add it to your .zshrc or .bashrc below
- ```export CHESSTERM_PUZZLES_DIR="[output from pwd]"```
## Step 3
- Run the program! ```chessTerm -elo 1200``` (or any elo you want to play)

## Outputs
<img width="1370" alt="Screenshot 2024-07-06 at 11 33 49 AM" src="https://github.com/MilesJuddPorter/chessTerm/assets/13202373/48fd0067-9ee1-456a-8b0e-c1e3ba218cdd">
<img width="935" alt="Screenshot 2024-07-06 at 11 34 15 AM" src="https://github.com/MilesJuddPorter/chessTerm/assets/13202373/7039108f-5371-4286-a41b-ca593a130c07">


## Other Information
- I just made this as my first go project to learn, feel free to do anything with it.
- This project uses lichess' open sourced puzzles
- Next steps are to add a player elo that will change as they play more puzzles, and to not show puzzles the player has already completed
