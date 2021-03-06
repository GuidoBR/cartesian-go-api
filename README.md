Cartesian API
=============

It's an API server in [go](https://golang.org/) that deals with a series of points represented as (x,y) coordinates on a simple 2-dimensional plane. Take a look at https://en.wikipedia.org/wiki/Cartesian_coordinate_system if you need a refresher on this concept.

`GET /api/points?x={d}&y={y}&distance={distance}`

- `x` integer (required). This represents the `x` coordinate of the search origin.
- `y` integer (required). This represents the `y` coordinate of the search origin.
- `distance` integer (required). This represents the Manhattan distance; points within `distance` from `x` and `y` are returned, points outside are filtered out.

Returns a JSON list of points that are within `distance` from `x,y`, using the Manhattan distance method. The points are returned in order of increasing distance from the search origin.

The Manhattan distance is measured "block-wise", as the distance in blocks between any two points in the plane (e.g. 2 blocks down and 3 blocks over for a total of 5 blocks). It is defined as the sum of the horizontal and vertical distances between points on a grid. Formally, where `p1 = (x1, y1)` and `p2 = (x2, y2)`, `distance(p1,p2) = |x1-x2| + |y1-y2|`.

On startup, the API server read a list of points from `data/points.json`.

## Run

```
go build && ./cartesian-go-api

curl -s http://127.0.0.1:8000/api/points\?x\=10\&y\=20\&distance\=20
```

## Tests

```
go test
```