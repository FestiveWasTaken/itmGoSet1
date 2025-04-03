// Placeholder
package main

import (
	"strings"
)

// Relationship status
//
// Let's pretend that you are building a new app with social media functionality.
// Users can have relationships with other users.
//
// The two guidelines for describing relationships are:
// 1. Any user can follow any other user.
// 2. If two users follow each other, they are considered friends.
//
// This function describes the relationship that two users have with each other.
//
// Please see the sample data for examples of `socialGraph`.
//
// Params:
// - fromMember, the subject member
// - toMember, the object member
// - socialGraph, the relationship data
//
// Returns:
// - "follower" if fromMember follows toMember; "followed by" if fromMember is followed by toMember; "friends" if fromMember and toMember follow each other; "no relationship otherwise."
func relationshipStatus(fromMember string, toMember string, socialGraph map[string]map[string]interface{}) string {
	//if fromMember follows toMember
	fromFollowing := socialGraph[fromMember]["following"].(string)
	toFollowing := socialGraph[toMember]["following"].(string)

	fromFollowsTo := false
	toFollowsFrom := false

	//if fromMember follows toMember
	if fromFollowing != "" {
		for _, user := range strings.Split(fromFollowing, ",") {
			if user == toMember {
				fromFollowsTo = true
				break
			}
		}
	}

	//if toMember follows fromMember
	if toFollowing != "" {
		for _, user := range strings.Split(toFollowing, ",") {
			if user == fromMember {
				toFollowsFrom = true
				break
			}
		}
	}

	//relationship status
	if fromFollowsTo && toFollowsFrom {
		return "friends"
	} else if fromFollowsTo {
		return "follower"
	} else if toFollowsFrom {
		return "followed by"
	}
	return "no relationship"
}

// Tic tac toe
//
// Tic Tac Toe is a common paper-and-pencil game.
// Players must attempt to draw a line of their symbol across a grid.
// The player that does this first is considered the winner.
//
// This function evaluates a Tic Tac Toe game board and returns the winner.
//
// Please see the sample data for examples of `board`.
//
// Params:
// - board, the representation of the Tic Tac Toe board as a square slice of slices of strings. The size of the slice will range between 3x3 to 6x6. The board will never have more than 1 winner. There will only ever be 2 unique symbols at the same time.
//
// Returns:
// - the symbol of the winner, or "NO WINNER" if there is no winner.
func ticTacToe(board [][]string) string {
	n := len(board)

	// Check rows
	for i := 0; i < n; i++ {
		if board[i][0] != "" && allSame(board[i]) {
			return board[i][0]
		}
	}

	// Check columns
	for j := 0; j < n; j++ {
		if board[0][j] != "" {
			winner := true
			for i := 1; i < n; i++ {
				if board[i][j] != board[0][j] {
					winner = false
					break
				}
			}
			if winner {
				return board[0][j]
			}
		}
	}

	// Check diagonal
	if board[0][0] != "" {
		winner := true
		for i := 1; i < n; i++ {
			if board[i][i] != board[0][0] {
				winner = false
				break
			}
		}
		if winner {
			return board[0][0]
		}
	}

	// Check anti-diagonal
	if board[0][n-1] != "" {
		winner := true
		for i := 1; i < n; i++ {
			if board[i][n-1-i] != board[0][n-1] {
				winner = false
				break
			}
		}
		if winner {
			return board[0][n-1]
		}
	}

	return "NO WINNER"
}

// Helper function for ticTacToe
func allSame(row []string) bool {
	for i := 1; i < len(row); i++ {
		if row[i] != row[0] {
			return false
		}
	}
	return true
}

// ETA
//
// A shuttle van service is tasked to travel one way along a predefined circular route.
// The route is divided into several legs between stops.
// The route is fully connected to itself.
//
// This function returns how long it will take the shuttle to arrive at a stop after leaving anothe rstop.
//
// Please see the sample data for examples of `routeMap`.
//
// Params:
// - firstStop, the stop that the shuttle will leave
// - secondStop, the stop that the shuttle will arrive at
// - routeMap, the data describing the routes
//
// Returns:
// - the time that it will take the shuttle to travel from firstStop to secondStop

// assume shuttle cant go backwards must loop around
func eta(firstStop string, secondStop string, routeMap map[string]map[string]int) int {
	// Create a map of all stops and their connections
	stops := make(map[string]map[string]int)

	// make graph from routeMap
	for route, data := range routeMap {
		stops1 := strings.Split(route, ",")
		from, to := stops1[0], stops1[1]
		time := data["travel_time_mins"]

		// maps
		if stops[from] == nil {
			stops[from] = make(map[string]int)
		}
		if stops[to] == nil {
			stops[to] = make(map[string]int)
		}

		//connect the stops and save their time
		stops[from][to] = time
		stops[to][from] = time
	}

	//assuming circular route make shortest path
	current := firstStop
	totalTime := 0
	visited := make(map[string]bool)

	for current != secondStop {
		visited[current] = true
		nextStop := ""
		for stop, time := range stops[current] {
			if !visited[stop] {
				nextStop = stop
				totalTime += time
				break
			}
		}
		if nextStop == "" {
			break
		}
		current = nextStop
	}

	return totalTime
}
