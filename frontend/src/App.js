import React, { useState } from "react";
import { BrowserRouter as Router, Route, Routes, Link } from "react-router-dom";
import "./App.css";
import Home from "./Home";
import Login from "./Login";
import Courses from "./Courses";
import Results from "./Results";
import Details from "./Details";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/login" element={<Login />} />
        <Route path="/courses" element={<Courses/>}  />
        <Route path="/results" element={<Results/>}  />
        <Route path="/details" element={<Details/>}  />
      </Routes>
    </Router>
  );
}

export default App;


// export default App;
// function App() {
//   const [courses, setCourses] = useState([{"title":"blabla", "price":50}]);
//   const [query, setQuery] = useState(''); // State for current input value
//   const [search, setSearch] = useState(''); // State for confirmed search term

//   useEffect(() => {
//     console.log(`Fetching data from http://localhost:8080/courses/search?query=${search}`)
//     // Fetch data from the API based on the search term
//     fetch(`http://localhost:8080/courses/search?query=${search}`)
//       .then(response => response.json())
//       .then(data => setCourses(data.results))
//       .catch(error => console.error('Error fetching courses:', error));
//   }, [search]);

//   const handleSearchChange = (e) => {
//     console.log(`Current query: ${e.target.value}`)
//     setQuery(e.target.value); // Update the query state as the user types
//   };

//   const handleSearchSubmit = (e) => {
//     e.preventDefault(); // Prevent the default form submission
//     setSearch(query); // Update the search state with the current query value
//   };

//   return (
//     <div className="App">
//       <div className="SearchBar">
//         <form onSubmit={handleSearchSubmit}>
//           <input
//             type="text"
//             placeholder="Search for courses..."
//             value={query}
//             onChange={handleSearchChange}
//           />
//         </form>
//       </div>
//       <div className="Courses">
//         {courses != null ? (
//           courses.map(course => (
//             <div key={course.id} className="Course">
//               <img src={course.image_url} alt={course.title} className="Course-image" />
//               <div className="Course-details">
//                 <h1 className="Course-title">{course.title}</h1>
//                 <p className="Course-description">{course.description}</p>
//                 <p className="Course-price">{course.price}</p>
//                 <p className="Course-category"><strong>{course.category}</strong></p>
//               </div>
//             </div>
//           ))
//         ) : (
//           <p>Loading courses...</p>
//         )}
//       </div>
//     </div>
//   );
// }

