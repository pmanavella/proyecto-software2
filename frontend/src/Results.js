import React, { useEffect, useState} from 'react'
//import { useNavigate, useParams } from 'react-router-dom';
import './Results.css';
import { useParams } from 'react-router-dom';
import Cookies from "universal-cookie";

const Cookie = new Cookies();

function goto(path){
    window.location = window.location.origin + path
}

const Results = () => {
    const[course, setCourse] =useState({});
    const { id } = useParams();
    // const navigate = useNavigate();

    // useEffect(() => {
    //     fetch(`http://localhost:8081/course/${id}`)
    //         .then(response => response.json())
    //         .then(data => setCourse(data));
    // }, [id]);


    // const handleViewDetails = () => {
    //     navigate(`/details/${id}`);
    // };

    const showCourses= () =>{
        return(
            <div>
                <div className="course-banner-small">CURSOS</div>
                <div class="course-title">
                    {course.title}
                </div>
                <div class="course-info">
                    <div>
                        <h4>Descripcion:</h4>
                        <p class="course-description">{course.description}</p>
                    </div>
                        <h4>Puntos:</h4>
                        <p class="course-points">{course.points}</p>
                    <div>
                        <div class="course-image">
                            <img src={course.image_url} alt={course.title}></img>
                        </div>
                        <button className="button" onClick={() => goto("/details")}>ver detalle</button>
                        {/* <button type="submit" onClick={handleViewDetails}>ver detalle</button> */}
                    </div>
                </div>
            </div>
        )
    }

    return (
        <div>
            {showCourses()}
            <button className="white-button" onClick={() => goto("/")}>HOME 🏠</button>
        </div>
    )
}

export default Results;

/*
import React, { useState, useEffect } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import './Results.css';

function goto(path){
    window.location = window.location.origin + path
}

const Results = () => {
    const [course, setCourse] = useState({});
    const { id } = useParams();
    const navigate = useNavigate();

    useEffect(() => {
        fetch(`http://localhost:8081/course/${id}`)
            .then(response => response.json())
            .then(data => setCourse(data));
    }, [id]);

    const handleViewDetails = () => {
        navigate(`/details/${id}`);
    };

    const showCourses = () => {
        return (
            <div>
                <div className="course-banner-small">CURSOS</div>
                <div className="course-title">
                    {course.title}
                </div>
                <div className="course-info">
                    <div>
                        <h4>Descripcion:</h4>
                        <p className="course-description">{course.description}</p>
                    </div>
                    <div className="course-image">
                        <img src={course.image_url} alt={course.title}></img>
                    </div>
                    <div>
                        <h4>Puntos:</h4>
                        <p className="course-points">{course.points}</p>
                        <button type="submit" onClick={handleViewDetails}>ver detalle</button>
                    </div>
                    <div>
                        <p className="course-category">{course.category}</p>
                    </div>
                </div>
            </div>
        );
    };

    return (
        <div>
            {showCourses()}
        </div>
    );
};

export default Results;
*/
