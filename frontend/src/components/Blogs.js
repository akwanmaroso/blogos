import React, { Component } from 'react';
import { render } from 'react-dom';
import ReadMore from './ReadMore';

export default class Blogs extends Component {
    constructor() {
        super();
        this.state = {
            posts: []
        };
    }

    componentDidMount() {
        fetch('http://localhost:9000/posts')
            .then(res => res.json())
            .then(data => this.setState({
                posts: data
            }))
    }

    getWrapperWidth() {
        if (this.wrapper) {
            console.log('get wrapper width', window.getComputedStyle(this.wrapper).getPropertyValue('width'));
        } else {
            console.log('get wrapper - no wrapper');
        }
    }

    render() {
        return (
            <div className="container" ref={node => this.wrapper = node}>
                <br />
                {
                    this.state.posts.map(post => {
                        return (
                            <div key={post.id}>
                                <h2>{post.title}</h2>
                                <small>
                                    author {post.author.nickname} - created at {new Intl.DateTimeFormat('en-US').format(new Date(post.created_at))} - updated at {new Intl.DateTimeFormat('en-US').format(new Date(post.updated_at))}
                                </small>
                                <ReadMore
                                    text={post.content}
                                    numberOfLines={4}
                                    lineHeight={1.4}
                                    showLessButton={true}
                                    onContentChange={this.getWrapperWidth}
                                />
                                <hr />
                            </div>
                        )
                    })
                }
            </div>
        );
    }
}

