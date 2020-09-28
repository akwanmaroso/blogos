import React, { useState, useEffect } from 'react';
import './App.css';

const App = () => {
    const [posts, setPosts] = useState([])

    useEffect(() => {
        fetch("http://localhost:9000/posts")
            .then(res => res.json())
            .then(data => setPosts(data))
            .catch(e => console.log(e))
    }, [])
    return (
        <div className="App">
            <h1>Blog</h1>
            {
                posts.map(p => {
                    return (
                        <>
                            <div key={p.id}><h2>{p.title}</h2></div>
                            <small>
                                author {p.author.nickname} - created at {new Intl.DateTimeFormat('en-US').format(new Date(p.created_at))} - updated at {new Intl.DateTimeFormat('en-US').format(new Date(p.updated_at))}
                            </small>
                            <hr />
                            <p>
                                {p.content}
                            </p>
                        </>
                    )
                })
            }
        </div>
    )
}

export default App;