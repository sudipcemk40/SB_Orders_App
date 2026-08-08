import { useEffect, useState } from "react";
import axios from "axios";

function App() {
  const [dashboard, setDashboard] = useState(null);
const [circulars, setCirculars] =useState([]);
const[searchNo, setSearchNo] =useState("");
const[searchResult, setSearchResult] =useState(null);
const[selectedYear, setSelectedYear] =useState("");

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
function searchByYear() {
  if (!selectedYear) {
    alert("Please select a year");
    return;
  }

  axios
    .get(`http://localhost:8080/year?year=${selectedYear}`)
    .then((res) => {
      console.log("Year result:", res.data);
      setCirculars(res.data);
    })
    .catch((err) => {
      console.log(err);
      alert("Unable to load circulars");
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
{/* Search By Year */}
<div className="row mb-4">

  <div className="col-md-8">
    <select
      className="form-select"
      value={selectedYear}
      onChange={(e) => setSelectedYear(e.target.value)}
    >
      <option value="">Select Year</option>
      <option value="2026">2026</option>
      <option value="2025">2025</option>
      <option value="2024">2024</option>
      <option value="2023">2023</option>
      <option value="2022">2022</option>
      <option value="2021">2021</option>
      <option value="2020">2020</option>
      <option value="2019">2019</option>
      <option value="2018">2018</option>
      <option value="2017">2017</option>
      <option value="2016">2016</option>
      <option value="2015">2015</option>
      <option value="2014">2014</option>
      <option value="2013">2013</option>
      <option value="2012">2012</option>
      <option value="2011">2011</option>
      <option value="2010">2010</option>
    </select>
  </div>

  <div className="col-md-4">
    <button
      className="btn btn-primary w-100"
      onClick={searchByYear}
    >
      Search by Year
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
  {Array.isArray(circulars) && circulars.length > 0 ? (
    circulars.map((c, index) => (
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
    ))
  ) : (
    <tr>
      <td colSpan="2" className="text-center">
        No circulars found
      </td>
    </tr>
  )}
</tbody>
</table>
    </table>

  </div>
);
}

export default App;