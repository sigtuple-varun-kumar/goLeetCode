/*
In a special ranking system, each voter gives a rank from highest to lowest to all teams participating in the competition.

The ordering of teams is decided by who received the most position-one votes. If two or more teams tie in the first position, we consider the second position to resolve the conflict, if they tie again, we continue this process until the ties are resolved. If two or more teams are still tied after considering all positions, we rank them alphabetically based on their team letter.

You are given an array of strings votes which is the votes of all voters in the ranking systems. Sort all teams according to the ranking system described above.

Return a string of all teams sorted by the ranking system.
Example 1:

Input: votes = ["ABC","ACB","ABC","ACB","ACB"]
Output: "ACB"
Explanation:
Team A was ranked first place by 5 voters. No other team was voted as first place, so team A is the first team.
Team B was ranked second by 2 voters and ranked third by 3 voters.
Team C was ranked second by 3 voters and ranked third by 2 voters.
As most of the voters ranked C second, team C is the second team, and team B is the third.


rankTeams determines the ranking of teams based on the given votes.

Parameters:
- votes: A slice of strings where each string represents a ranking of teams by a voter.

Returns:
- A string representing the final ranking of teams based on the votes.

Note:
- If there is only one vote, the function directly returns that vote as the ranking.
*/

package leetcode

import "sort"

func rankTeams(votes []string) string {
	if len(votes) == 1 {
		return votes[0]
	}
	teams := votes[0]
	teamRunes := []rune(teams)
	teamCount := len(teams)
	countsPerTeam := make([][26]int, teamCount)

	for _, vote := range votes {
		for i, team := range vote {
			countsPerTeam[i][team-'A']++
		}
	}
	sort.Slice(teamRunes, func(i, j int) bool {
		teamI, teamJ := teamRunes[i]-'A', teamRunes[j]-'A'
		for pos := 0; pos < teamCount; pos++ {
			voteI := countsPerTeam[pos][teamI]
			voteJ := countsPerTeam[pos][teamJ]
			if voteI != voteJ {
				return voteI > voteJ
			}
		}
		return teamRunes[i] < teamRunes[j]
	})
	return string(teamRunes)
}
