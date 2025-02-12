SESSION_NAME="algo-schedule"

tmux new-session -d -s $SESSION_NAME -n editor
tmux send-keys -t $SESSION_NAME "nvim ." C-m

tmux new-window -t $SESSION_NAME -n styles
tmux send-keys -t $SESSION_NAME "./tailwindcss -m -i ./static/global.css -o ./static/global.rendered.css && templ generate && go run . -dev" C-m

tmux new-window -t $SESSION_NAME -n server

tmux select-window -t $SESSION_NAME:1
tmux attach -t $SESSION_NAME
