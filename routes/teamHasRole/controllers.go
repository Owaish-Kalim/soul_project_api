package teamHasRole

import (
	"encoding/json"
	"fmt"
	"net/http"
	"soul_api/config"
	Shared "soul_api/routes"
	"strconv"
	// "github.com/gemcook/pagination-go"
)

func TeamHasRole(w http.ResponseWriter, r *http.Request) ([]TeamRole, Shared.ErrorMsg) {
	r.ParseForm()
	var response []TeamRole
	q := &query{}
	limit := r.Form.Get("limit")
	if limit != "" {
		if err := Shared.ParseInt(r.Form.Get("limit"), &q.Limit); err != nil {
			return response, Shared.ErrorMsg{Message: "parseerr"}
		}
	} else {
		q.Limit = 10
	}
	page := r.Form.Get("page")
	if page != "" {
		if err := Shared.ParseInt(r.Form.Get("page"), &q.Page); err != nil {
			return response, Shared.ErrorMsg{Message: "parseerr"}
		}
		q.Page = q.Page - 1
	} else {
		q.Page = 0
	}
	teamid := r.Form.Get("teamid")
	if teamid != "" {
		if err := Shared.ParseInt(r.Form.Get("teamid"), &q.TeamId); err != nil {
			return response, Shared.ErrorMsg{Message: "parseerr"}
		}
	}
	q.Status = r.Form.Get("status")
	q.FirstName = r.Form.Get("firstname")
	q.LastName = r.Form.Get("lastname")

	fmt.Println(q)
	offset := q.Limit * q.Page

	var teamRoles []TeamRole
	fmt.Println(12)
	sqlStatement := `SELECT ("Team_Has_Role_Id"),("FirstName"), ("LastName"), ("Team_Id"),("Status"),("CreatedAt"),("UpdatedAt") FROM slh_team_has_role 
	WHERE ("Status") LIKE ''|| $1 ||'%' 
	AND ("FirstName") LIKE ''|| $2 ||'%' 
	AND ("LastName") LIKE ''|| $3 ||'%' 
	ORDER BY ("CreatedAt") DESC LIMIT $4 OFFSET $5`
	rows, err := config.Db.Query(sqlStatement, q.Status, q.FirstName, q.LastName, q.Limit, offset)
	fmt.Println(13)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return response, Shared.ErrorMsg{Message: "Internal Server Error."}
	}

	for rows.Next() {
		var team = TeamRole{}
		rows.Scan(&team.Team_Has_Role_Id, &team.FirstName, &team.LastName, &team.TeamId, &team.Status, &team.CreatedAt, &team.UpdatedAt)
		teamRoles = append(teamRoles, team)
	}

	sqlStatement = `SELECT COUNT(*) FROM slh_team_has_role 
	WHERE ("Status") LIKE ''|| $1 ||'%'
	AND ("FirstName") LIKE ''|| $2 ||'%' 
	AND ("LastName") LIKE ''|| $3 ||'%'`
	cntRow := config.Db.QueryRow(sqlStatement, q.Status, q.FirstName, q.LastName)
	fmt.Println(14)
	cnt := 0
	err = cntRow.Scan(&cnt)
	if err != nil {
		panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		return teamRoles, Shared.ErrorMsg{Message: "Internal Server Error."}
	}

	w.Header().Set("Total-Count", strconv.Itoa(cnt))
	totalPages := cnt / q.Limit
	if cnt%q.Limit != 0 {
		totalPages = totalPages + 1
	}

	w.Header().Set("Total-Pages", strconv.Itoa(totalPages))

	fmt.Println(cnt)
	fmt.Println(15)
	return teamRoles, Shared.ErrorMsg{Message: ""}
}

func Role_Team(w http.ResponseWriter, r *http.Request) (Roles, ErrorMessage) {
	r.ParseForm()
	roles := Roles{}
	err := json.NewDecoder(r.Body).Decode(&roles)
	if err != nil {
		panic(err)
	}

	roles.Status = true

	sqlStatement := `
	INSERT INTO slh_roles ("Role_Name", "Role_Status") 
	VALUES ($1, $2)
	RETURNING ("Role_Id")`

	roles.Role_Id = 0
	err = config.Db.QueryRow(sqlStatement, roles.Role_Name, roles.Status).Scan(&roles.Role_Id)
	if err != nil {
		return Roles{}, ErrorMessage{Message: "Role_Name already Registered"}
	}

	return roles, ErrorMessage{}
}

func TeamHasRoleUpdate(w http.ResponseWriter, r *http.Request) (RoleUp, ErrorMessage) {
	fmt.Println(4)
	r.ParseForm()
	roles := RoleUp{}
	err := json.NewDecoder(r.Body).Decode(&roles)
	if err != nil {
		panic(err)
	}

	sqlStatement := `SELECT ("Role_Id") FROM slh_roles WHERE ("Role_Name")=$1;`
	row := config.Db.QueryRow(sqlStatement, roles.Role_Name)
	err = row.Scan(&roles.Role_Id)
	if err != nil {
		fmt.Println(1)
		panic(err)
	}

	sqlStatement = ` UPDATE slh_team_has_role SET "Team_Has_Role_Id" = $1  WHERE ("Team_Id") = $2`

	_, err = config.Db.Exec(sqlStatement, roles.Role_Id, roles.TeamId)
	if err != nil {
		fmt.Println(2)
		panic(err)
	}

	sqlStatement = ` UPDATE slh_teams SET "Role" = $1  WHERE ("TeamId") = $2`

	_, err = config.Db.Exec(sqlStatement, roles.Role_Name, roles.TeamId)
	if err != nil {
		fmt.Println(3)
		panic(err)
	}

	return roles, ErrorMessage{}
}
