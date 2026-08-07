import { useEffect, useState } from "react";
import axios from "axios";

function App() {
  const [dashboard, setDashboard] = useState(null);
const [circulars, setCirculars] =useState([]);
const[searchNo, setSearchNo] =useState("");
const[searchResult, setSearchResult] =useState(null);

  useEffect(() => {
    axios.get("http://localhost:8080/dashboard")
      .then((res) => setDashboard(res.data))
      .catch((err) => console.log(err));
axios.get("http://localhost:8080/circulars")
      .then ((res) => {
console.log("Circulars:", res.data);
setCirculars(res.data);
})
.catch((err) => console.log(err));
  }, []);
console.log(circulars);
function searchCircular() {
  axios.get(
`http://localhost:8080/search?no=${searchNo.trim()}`
)
    .then((res) => {
      setSearchResult(res.data);
    })
    .catch((err) => {
      alert("Circular not found");
      console.log(err);
    });
}
return (
  <div className="container mt-5">

    <div className="bg-primary text-white p-4 rounded mb-4">
      <h2>SB Orders Portal</h2>
      <p>Department of Posts</p>
    </div>

    {/* Dashboard Cards */}
    {dashboard && (
      <div className="row">
        ...
      </div>
    )}

    {/* Search Box */}
    <div className="row mt-4 mb-4">
      <div className="col-md-8">
        <input
          type="text"
          className="form-control"
          placeholder="Enter Circular Number"
          value={searchNo}
          onChange={(e) => setSearchNo(e.target.value)}
        />
      </div>

      <div className="col-md-4">
        <button
          className="btn btn-primary w-100"
          onClick={searchCircular}
        >
          Search
        </button>
      </div>
    </div>

    {/* ADD THIS HERE */}
    {searchResult && (
      <div className="alert alert-success">
        <h5>{searchResult.circularNo}</h5>

        <a
          href={`http://localhost:8080/pdf?no=${searchResult.circularNo}`}
          target="_blank"
          rel="noreferrer"
        >
          📄 Open PDF
        </a>
      </div>
    )}

    {/* All Circulars Table */}
    <h3 className="mt-5">All Circulars</h3>

    <table className="table table-bordered table-striped">
      <table className="table table-bordered table-striped">
  <thead>
    <tr>
      <th>Circular No</th>
      <th>PDF</th>
    </tr>
  </thead>

  <tbody>
    {circulars.map((c, index) => (
      <tr key={index}>
        <td>{c.circularNo}</td>
        <td>
          <a
            href={`http://localhost:8080/pdf?no=${c.circularNo}`}
            target="_blank"
            rel="noreferrer"
          >
            📄 Open PDF
          </a>
        </td>
      </tr>
    ))}
  </tbody>
</table>
    </table>

  </div>
);
}

export default App;