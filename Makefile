RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
$(eval $(RUN_ARGS):;@:)

help: 
	@echo "Commands:"
	@echo " - help: 	show this help"
	@echo " - new:		create a new file for today"

setup: 
	@touch session.cookie
	@mkdir -p inputs
	@make help

new: 
	@go run src/make/new_file.go

solve:
	@go run src/solutions/day$(RUN_ARGS)/solution.go