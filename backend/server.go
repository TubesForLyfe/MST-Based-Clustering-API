package main

import (
	"fmt"
	"log"
	"os"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"database/sql"
	"strconv"
	"math"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Point struct {
	x int
	y int
}

type MinimumSpanningTree struct {
	p1 Point
	p2 Point
}

func getEnv(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func openDatabase() *sql.DB {
	// Open database connection.
	db, err := sql.Open("mysql",  getEnv("DATABASE_USERNAME")+":"+ getEnv("DATABASE_PASSWORD")+"@tcp("+ getEnv("DATABASE_PORT")+")/"+ getEnv("DATABASE_NAME"))

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	return db
}

func isSame(p1 Point, p2 Point) bool {
	return p1.x == p2.x && p1.y == p2.y
}

func isIn(p Point, listPoint []Point) bool {
	var i int
	for i = 0; i < len(listPoint); i++ {
		if (isSame(p, listPoint[i])) {
			return true
		}
	}
	return false
}

func getDistance(p1 Point, p2 Point) float64 {
	return math.Sqrt(math.Pow(float64(p2.x - p1.x), 2) + math.Pow(float64(p2.y - p1.y), 2))
}

func isEdge(p1 Point, p2 Point, MST []MinimumSpanningTree) bool {
	var i int
	for i = 0; i < len(MST); i++ {
		if (isSame(MST[i].p1, p1) && isSame(MST[i].p2, p2)) {
			return true
		}
	}
	return false
}

func deleteFirstPoint(listPoint []Point) []Point {
	var i int
	var result []Point
	
	for i = 1; i < len(listPoint); i++ {
		result = append(result, listPoint[i])
	}
	return result
}

func deletePointByValue(p Point, listPoint []Point) []Point {
	var i int
	var result []Point
	
	for i = 0; i < len(listPoint); i++ {
		if (!isSame(p, listPoint[i])) {
			result = append(result, listPoint[i])
		}
	}
	return result
}

func deleteLastSpanningTree(MST []MinimumSpanningTree) []MinimumSpanningTree {
	var i int
	var result []MinimumSpanningTree

	for i = 0; i < len(MST) - 1; i++ {
		result = append(result, MST[i])
	}
	return result
}

func isSirkuler(p1 Point, p2 Point, MST []MinimumSpanningTree) bool {
	var i int
	var queue []Point
	var processing Point
	var processed []Point

	queue = append(queue, p1)
	for len(queue) != 0 {
		processing = queue[0]
		processed = append(processed, processing)
		queue = deleteFirstPoint(queue)
		for i = 0; i < len(MST); i++ {
			if (isSame(MST[i].p1, processing)) {
				if (isSame(MST[i].p2, p2)) {
					return true
				} else {
					if (!isIn(MST[i].p2, processed)) {
						queue = append(queue, MST[i].p2)
					}
				}
			} else {
				if (isSame(MST[i].p2, processing)) {
					if (isSame(MST[i].p1, p2)) {
						return true
					} else {
						if (!isIn(MST[i].p1, processed)) {
							queue = append(queue, MST[i].p1)
						}
					}
				}
			}
		}
	}
	return false
}

func KruskalAlgorithm(listPoint []Point) []MinimumSpanningTree {
	var i, j int
	var min float64
	var result []MinimumSpanningTree
	
	for len(result) < len(listPoint) - 1 {
		var tree MinimumSpanningTree

		min = -1
		for i = 0; i < len(listPoint); i++ {
			for j = i + 1; j < len(listPoint); j++ {
				if !isEdge(listPoint[i], listPoint[j], result) {
					if !isSirkuler(listPoint[i], listPoint[j], result) {
						if getDistance(listPoint[i], listPoint[j]) < min || min < 0 {
							tree.p1 = listPoint[i]
							tree.p2 = listPoint[j]
							min = getDistance(listPoint[i], listPoint[j])
						}
					}
				}
			}
		}
		result = append(result, tree)
	}
	return result
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func CreateMST(res http.ResponseWriter, req *http.Request) {
	setupResponse(&res, req)
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}
	pointData := make(map[string][]map[string]string)
	json.Unmarshal(body, &pointData)

	if len(pointData["data"]) != 0 {
		var point Point
		var listPoint []Point
		var i int
		var MST []MinimumSpanningTree

		for i = 0; i < len(pointData["data"]); i++ {
			point.x, _ = strconv.Atoi(pointData["data"][i]["x"])
			point.y, _ = strconv.Atoi(pointData["data"][i]["y"])
			listPoint = append(listPoint, point)
		}
		MST = KruskalAlgorithm(listPoint)
		MSTResult := []map[string][]map[string]int{}
		for i = 0; i < len(MST); i++ {
			MSTMap := make(map[string][]map[string]int)
			MSTPoint := []map[string]int{}
			resultPoint := make(map[string]int)

			resultPoint["x"] = MST[i].p1.x
			resultPoint["y"] = MST[i].p1.y
			MSTPoint = append(MSTPoint, resultPoint)
			resultPoint = make(map[string]int)
			resultPoint["x"] = MST[i].p2.x
			resultPoint["y"] = MST[i].p2.y
			MSTPoint = append(MSTPoint, resultPoint)
			MSTMap["point"] = MSTPoint
			MSTResult = append(MSTResult, MSTMap)
		}
		response, _ := json.Marshal(MSTResult)
		res.Write(response)
	}
}

func CreateCluster(res http.ResponseWriter, req *http.Request) {
	setupResponse(&res, req)
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}
	data := make(map[string]string)
	json.Unmarshal(body, &data)
	cluster_amount, err := strconv.Atoi(data["amount"])
	filename := data["filename"]

	if cluster_amount > 0 {
		var point Point
		var i, j int
		var ST MinimumSpanningTree
		var MST []MinimumSpanningTree

		data := make(map[string][]map[string][]map[string]int)
		json.Unmarshal(body, &data)
		MSTRequest := data["MST"]
		if (cluster_amount <= len(MSTRequest) + 1) {
			cluster_list := []Point{}
			for i = 0; i < len(MSTRequest); i++ {
				point.x = MSTRequest[i]["point"][0]["x"]
				point.y = MSTRequest[i]["point"][0]["y"]
				if (!isIn(point, cluster_list)) {
					cluster_list = append(cluster_list, point)
				}
				ST.p1 = point
				point.x = MSTRequest[i]["point"][1]["x"]
				point.y = MSTRequest[i]["point"][1]["y"]
				if (!isIn(point, cluster_list)) {
					cluster_list = append(cluster_list, point)
				}
				ST.p2 = point
				MST = append(MST, ST)
			}

			amount := cluster_amount
			for amount > 1 {
				MST = deleteLastSpanningTree(MST)
				amount--
			}

			cluster_point := [][]Point{}
			for len(cluster_point) < cluster_amount {
				cluster := []Point{}
				cluster = append(cluster, cluster_list[0])
				queue := cluster
				cluster_list = deleteFirstPoint(cluster_list)
				for len(queue) != 0 {
					processing := queue[0]
					queue = deleteFirstPoint(queue)
					for i = 0; i < len(MST); i++ {
						if (isSame(processing, MST[i].p1)) {
							if (isIn(MST[i].p2, cluster_list)) {
								cluster = append(cluster, MST[i].p2)
								queue = append(cluster, MST[i].p2)
								cluster_list = deletePointByValue(MST[i].p2, cluster_list)
							}
						} else {
							if (isSame(processing, MST[i].p2)) {
								if (isIn(MST[i].p1, cluster_list)) {
									cluster = append(cluster, MST[i].p1)
									queue = append(cluster, MST[i].p1)
									cluster_list = deletePointByValue(MST[i].p1, cluster_list)
								}
							}
						}
					}
				}
				cluster_point = append(cluster_point, cluster)
			}

			cluster_result := [][]map[string]int{}
			for i = 0; i < len(cluster_point); i++ {
				cluster_data := []map[string]int{}
				for j = 0; j < len(cluster_point[i]); j++ {
					cluster_data_row := make(map[string]int)
					cluster_data_row["x"] = cluster_point[i][j].x
					cluster_data_row["y"] = cluster_point[i][j].y
					cluster_data = append(cluster_data, cluster_data_row)
				}
				cluster_result = append(cluster_result, cluster_data)
			}
			response, _ := json.Marshal(cluster_result)
			db := openDatabase()
			db_result, _ := db.Query("INSERT INTO cluster_log (filename, cluster_amount, cluster_result) VALUES ('" + filename + "', " + strconv.Itoa(cluster_amount) + ", '" + string(response) + "')")
			defer db_result.Close()
			defer db.Close()
			res.Write(response)
		}
	}
}

func GetClusterLog(res http.ResponseWriter, req *http.Request) {

}

func main() {
	http.HandleFunc("/create-MST", CreateMST)
	http.HandleFunc("/create-cluster", CreateCluster)
	http.HandleFunc("/get-cluster-log", GetClusterLog)

	fmt.Println("Starting server at port " + getEnv("BACKEND_PORT"))
	if err := http.ListenAndServe(":"+ getEnv("BACKEND_PORT"), nil); err != nil {
		log.Fatal(err)
	}
}