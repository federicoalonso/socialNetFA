import React, { useState, useEffect } from 'react';
import { Image } from 'react-bootstrap';
import { map } from 'lodash';

import './ListTweets.scss';

export default function ListTweets(props) {
    const { tweets } = props;


    return (
        <div className='list-tweet'>
            {map(tweets, (tweet, index) => (
                <Tweet tweet={tweet} key={index} />
            ))}
        </div>
    );
}

function Tweet(props) {
    const { tweet } = props;

    return <h2>{tweet.mensaje}</h2>
}