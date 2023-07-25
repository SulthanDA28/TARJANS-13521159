package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"math"
)
type Graph struct {
	node map[string][]string
	urut []string
}

func NewGraph() *Graph {
	return &Graph{node: make(map[string][]string)}
}
func (g *Graph) addEdge(u, v string) {
	g.node[u] = append(g.node[u], v)
}

func stringToMatrix(input string) [][]string {
	var matrix [][]string
	hasil := strings.Split(input, "\n")
	for _, v := range hasil {
		matrix = append(matrix, strings.Split(v, " "))
	}
	return matrix
}
func (g *Graph) PrintEdges() {
	for node, neighbors := range g.node {
		fmt.Println(node, ":", neighbors)
	}
}

func matrixtoGraph(matrix [][]string) Graph {
	graph := NewGraph()
	for i := 0; i < len(matrix); i++ {
		graph.addEdge(matrix[i][0], matrix[i][1])
	}
	graph.urut = matrixtoarr(matrix)
	return *graph
}
func stringToGraph(input string) Graph {
	matrix := stringToMatrix(input)
	graph := matrixtoGraph(matrix)
	return graph
}

func minint(i1, i2 int) int {
	if i1 < i2 {
		return i1
	}
	return i2
}

func matrixtoarr(matrix [][]string) []string {
	var arr []string
	unique := make(map[string]bool)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			unique[matrix[i][j]] = false
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if(unique[matrix[i][j]] == false){
				arr = append(arr, matrix[i][j])
				unique[matrix[i][j]] = true
			}
		}
	}

	return arr
}

func (g *Graph) TarjanSCC() [][]string {
	var (
		indeks int
		stack []string
		indices = make(map[string]int)
		minim = make(map[string]int)
		onStack = make(map[string]bool)
		hasil [][]string
		findSCC func(string)
	)
	findSCC = func(v string) {
		indices[v] = indeks
		minim[v] = indeks
		indeks++
		stack = append(stack, v)
		onStack[v] = true
		for _, neigbor := range g.node[v] {
			if indices[neigbor] < 0 {
				findSCC(neigbor)
				minim[v] = minint(minim[v], minim[neigbor])
			} else if onStack[neigbor] {
				minim[v] = minint(minim[v], indices[neigbor])
			}
		}
		if minim[v] == indices[v] {
			var scc []string
			for {
				pop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				onStack[pop] = false
				scc = append(scc, pop)
				if pop == v {
					break
				}
			}
			hasil = append(hasil, scc)
		}
	}
	for v := range g.node {
		indices[v] = -1
	}
	for v := range g.node {
		if indices[v] < 0 {
			findSCC(v)
		}
	}
	return hasil
}

func (g *Graph) TarjanBridge() [][]string {
	var(
		visited = make(map[string]bool)
		indices = make(map[string]float64)
		minim = make(map[string]float64)
		parent = make(map[string]string)
		hasil [][]string
		indeks int
		findBridge func(string)
	)
	findBridge = func(v string) {
		visited[v] = true
		indices[v] = float64(indeks)
		minim[v] = float64(indeks)
		indeks++
		for _, neighbor := range g.node[v] {
			if !visited[neighbor] {
				parent[neighbor] = v
				findBridge(neighbor)
				minim[v] = min(minim[v], minim[neighbor])
				if minim[neighbor] > indices[v] {
					hasil = append(hasil, []string{v, neighbor})
				}
			} else if neighbor != parent[v] {
				minim[v] = min(minim[v], indices[neighbor])
			}
		}
	}
	for v := range g.node {
		visited[v] = false
		indices[v] = math.Inf(1)
		minim[v] = math.Inf(1)
		parent[v] = ""
	}
	for v := range g.node {
		for v := range g.node {
			visited[v] = false
			indices[v] = math.Inf(1)
			minim[v] = math.Inf(1)
			parent[v] = ""
		}
		if !visited[v] {
			findBridge(v)
		}
	}
	return hasil
}


func min(i1, i2 float64) float64 {
	if i1 < i2 {
		return i1
	}
	return i2
}

type jsonhasil struct {
	Hasil string `json:"hasil"`
}

type jsonhasilakhir struct {
	Hasilscc []jsonhasil `json:"hasilscc"`
	Hasilbridge []jsonhasil `json:"hasilbridge"`
	Runtime string `json:"runtime"`
}

func (g *Graph) ubahhasilSCC() []jsonhasil{
	var hasilakhir string
	hasiljson := []jsonhasil{}
	hasilscc := g.TarjanSCC()
	for i := 0; i < len(hasilscc); i++ {
		if len(hasilscc[i]) == 1 {
			hasilakhir = hasilakhir + hasilscc[i][0]
		}else {
			for j := 0; j < len(hasilscc[i]); j++ {
				for k := 0; k < len(hasilscc[i]); k++ {
					for l := 0; l < len(g.node[hasilscc[i][k]]); l++ {
						if g.node[hasilscc[i][k]][l] == hasilscc[i][j] {
							hasilakhir = hasilakhir + hasilscc[i][k] + "->" + hasilscc[i][j] + " "
						}
						
					}
				}
			}
		}
		hasiljson = append(hasiljson, jsonhasil{Hasil: hasilakhir})
		hasilakhir = ""
	}
	return hasiljson
	
}
func (g *Graph) ubahhasilBridge() []jsonhasil{
	hasilakhir := []jsonhasil{}
	hilangduplikatmatrix := make(map[[2]string]bool)
	hasililangduplicat := g.TarjanBridge()
	for i := 0; i < len(hasililangduplicat); i++ {
		hilangduplikatmatrix[[2]string{hasililangduplicat[i][0], hasililangduplicat[i][1]}] = false
	}
	var hasilbridge [][]string
	for k := range hasililangduplicat {
		if hilangduplikatmatrix[[2]string{hasililangduplicat[k][0], hasililangduplicat[k][1]}] == false {
			hasilbridge = append(hasilbridge, hasililangduplicat[k])
			hilangduplikatmatrix[[2]string{hasililangduplicat[k][0], hasililangduplicat[k][1]}] = true
			hilangduplikatmatrix[[2]string{hasililangduplicat[k][1], hasililangduplicat[k][0]}] = true
		}
	}
	for i := 0; i < len(hasilbridge); i++ {
		hasilakhir = append(hasilakhir, jsonhasil{Hasil: hasilbridge[i][0] + "->" + hasilbridge[i][1]})
	}
	return hasilakhir
}
type jsoninput struct {
	Input string `json:"input"`
}

//------------------------------------------------
//---------------A P I----------------------------
//------------------------------------------------

func getHasil(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	start := time.Now()
	var input jsoninput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	graph := stringToGraph(input.Input)
	hasilbridge := graph.ubahhasilBridge()
	hasilscc := graph.ubahhasilSCC()
	end := time.Now()
	eksekusi := end.Sub(start)
	hasilakhir := jsonhasilakhir{Hasilscc: hasilscc, Hasilbridge: hasilbridge, Runtime: eksekusi.String()}
	c.IndentedJSON(http.StatusOK, hasilakhir)
}
func main() {
	router := gin.Default()
	router.POST("/hasil", getHasil)
	router.Use(cors.Default())
	router.Run("localhost:8080")
}

