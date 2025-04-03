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
	//determine wincons, diag, rows, column, anti-diag,
	//rows
	for i := 0; i < n; i++ {
		if board[i][0] != "" {
			allSame := true
			for j := 1; j < n; j++ {
				if board[i][j] == "" || board[i][j] != board[i][0] {
					allSame = false
					break
				}
			}
			if allSame {
				return board[i][0]
			}
		}
	}

	// columns
	for j := 0; j < n; j++ {
		if board[0][j] != "" {
			allSame := true
			for i := 1; i < n; i++ {
				if board[i][j] == "" || board[i][j] != board[0][j] {
					allSame = false
					break
				}
			}
			if allSame {
				return board[0][j]
			}
		}
	}

	// diagonal
	if board[0][0] != "" {
		allSame := true
		for i := 1; i < n; i++ {
			if board[i][i] == "" || board[i][i] != board[0][0] {
				allSame = false
				break
			}
		}
		if allSame {
			return board[0][0]
		}
	}

	// anti-diagonal
	if board[0][n-1] != "" {
		allSame := true
		for i := 1; i < n; i++ {
			if board[i][n-1-i] == "" || board[i][n-1-i] != board[0][n-1] {
				allSame = false
				break
			}
		}
		if allSame {
			return board[0][n-1]
		}
	}

	return "NO WINNER"
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

func eta(firstStop string, secondStop string, routeMap map[string]map[string]int) int {
	// map of all stops and their connections
	stops := make(map[string]map[string]int)

	// make graph from routeMap then reference graph for everything
	for route, data := range routeMap {
		stops1 := strings.Split(route, ",")
		from, to := stops1[0], stops1[1]
		time := data["travel_time_mins"]

		if stops[from] == nil {
			stops[from] = make(map[string]int)
		}
		if stops[to] == nil {
			stops[to] = make(map[string]int)
		}

		// clockwise only
		stops[from][to] = time
	}

	// in clockwise direction
	current := firstStop
	totalTime := 0
	visited := make(map[string]bool)

	for current != secondStop {
		visited[current] = true
		nextStop := ""
		minTime := -1

		// unvisited neighbor with mintime
		for stop, time := range stops[current] {
			if !visited[stop] && (minTime == -1 || time < minTime) {
				nextStop = stop
				minTime = time
			}
		}

		if nextStop == "" {
			//in the case where no unvisited neighbor is found, went wrong way
			//reset and try different path
			current = firstStop
			totalTime = 0
			visited = make(map[string]bool)
			continue
		}

		totalTime += minTime
		current = nextStop
	}

	return totalTime
}
