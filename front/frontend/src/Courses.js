import React, { useEffect, useState} from 'react'
import './Courses.css';
import { useParams } from 'react-router-dom';
import Cookies from "universal-cookie";

const Cookie = new Cookies();

async function getCourseById(id){
    return await fetch('http://localhost:8081/course/' + id, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    }
}).then(response => response.json())
}

function goto(path){
    window.location = window.location.origin + path
}

const Courses = () => {
    const[course, setCourse] =useState({});
    const { id } = useParams();

    if (!course.id_course) {
        getCourseById(Number(id)).then(response => {setCourse(response);})
    }
    const showCourses= () =>{
        return(
            <div>
                <div className="course-banner">MIS CURSOS</div>
                <div class="course-title">
                    {course.title}
                </div>
                <div class="course-info">
                    <div>
                        <h4>Descripcion:</h4>
                        <p class="course-description">{course.description}</p>
                    </div>
                    <div>
                        
                        <p class="course-category">{course.category}</p>
                    </div>
                    <div>
                        
                        <p class="course-instructor">{course.instructor}</p>
                    </div>
                    <div>
                        
                        <p class="course-duration">{course.duration}</p>
                    </div>
                    <div>
                        
                        <p class="course-capacity">{course.capacity}</p>
                    </div>
                    <div>
                        
                        <p class="course-points">{course.points}</p>
                    </div>
                    <div>
                        <div class="course-requirements-container">
                            <h4>Requisitos</h4>
                            <p class="course-requirements">{course.requirements}</p>
                        </div>
                        <div class="course-image">
                            <img src={course.image_url} alt={course.title}></img>
                        </div>
                    </div>
                </div>
            </div>
        )
    }

    return (
        <div>
            <button onClick={() => goto("/")}>HOME 🏠</button>
            {showCourses()}
        </div>
    )
}

export default Courses;
// import React, { useEffect, useState } from 'react';
// import './Courses.css';
// import { useParams } from 'react-router-dom';
// import Cookies from 'universal-cookie';

// const cookies = new Cookies();

// async function getCourseById(id) {
//   const response = await fetch(`http://localhost:8081/course/${id}`, {
//     method: 'GET',
//     headers: {
//       'Content-Type': 'application/json',
//     },
//   });
//   return response.json();
// }

// function goto(path) {
//   window.location = window.location.origin + path;
// }

// const Courses = () => {
//   const [course, setCourse] = useState({});
//   const { id } = useParams();

//   useEffect(() => {
//     if (!course.id_course) {
//       getCourseById(Number(id)).then((response) => {
//         setCourse(response);
//       });
//     }
//   }, [id, course.id_course]);

//   const showCourses = () => {
//     return (
//       <div>
//         <div className="course-banner-small">CURSOS</div>
//         <div className="course-title">{course.title}</div>
//         <div className="course-info">
//           <div>
//             <h4>Descripcion:</h4>
//             <p className="course-description">{course.description}</p>
//           </div>
//           <div>
//             <h4>Requisitos</h4>
//             <p className="course-requirements">{course.requirements}</p>
//           </div>
//           <div className="course-image">
//             <img src={course.image_url} alt={course.title}></img>
//           </div>
//         </div>
//       </div>
//     );
//   };

//   return (
//     <div>
//       <button onClick={() => goto("/")}>HOME 🏠</button>
//       {showCourses()}
//     </div>
//   );
// };

// export default Courses;

///// alert
// fectCourses Utiliza parámetros de consulta en la URL para especificar el tipo de búsqueda y el término de búsqueda.
///// alert


// import React, { useEffect, useState } from 'react';
// import './Courses.css';
// import { useParams, useLocation } from 'react-router-dom';
// import Cookies from 'universal-cookie';

// const cookies = new Cookies();

// async function fetchCourses(query) {
//   const response = await fetch(`http://localhost:8081/courses${query}`, {
//     method: 'GET',
//     headers: {
//       'Content-Type': 'application/json',
//     },
//   });
//   return response.json();
// }

// function goto(path) {
//   window.location = window.location.origin + path;
// }

// const Courses = () => {
//   const [courses, setCourses] = useState([]);
//   const [loading, setLoading] = useState(true);
//   const [error, setError] = useState(null);
//   const { id } = useParams();
//   const location = useLocation();

//   useEffect(() => {
//     const fetchCourseData = async () => {
//       try {
//         let query = '';
//         if (id) {
//           query = `/${id}`;
//         } else {
//           query = location.search;
//         }
//         const courseData = await fetchCourses(query);
//         setCourses(courseData);
//       } catch (err) {
//         setError(err);
//       } finally {
//         setLoading(false);
//       }
//     };

//     fetchCourseData();
//   }, [id, location.search]);

//   if (loading) {
//     return <div>Loading...</div>;
//   }

//   if (error) {
//     return <div>Error: {error.message}</div>;
//   }

//   const showCourses = () => {
//     return courses.map((course) => (
//       <div key={course.id_course}>
//         <div className="course-title">{course.title}</div>
//         <div className="course-description">{course.description}</div>
//         <h4>Requisitos</h4>
//         <p className="course-requirements">{course.requirements}</p>
//         <div className="course-image">
//           <img src={course.image_url} alt={course.title}></img>
//         </div>
//       </div>
//     ));
//   };

//   return (
//     <div>
//       <button onClick={() => goto("/")}>HOME 🏠</button>
//       {showCourses()}
//     </div>
//   );
// };

// export default Courses;
