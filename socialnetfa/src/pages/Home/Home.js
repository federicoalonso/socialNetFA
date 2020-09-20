import React from 'react';
import BasicLayout from "../../layout/BasicLayout";

import "./Home.scss";

export default function Home(props) {
    const {setRefreshChckLogin} = props;

    return (
            <BasicLayout className="home" setRefreshChckLogin={setRefreshChckLogin}>
                <h2>estamos en Home</h2>
            </BasicLayout>
    )
}
