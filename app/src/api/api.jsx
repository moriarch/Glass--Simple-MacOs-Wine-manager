const request = (path,props) => fetch(path,{
    ...props
}).then(res=>res.json())

const local = 'http://localhost:8181';

const apiRoutes = {
    main: local+'/api/main',
    changeFolder:local+'/api/main/changeFolder',
    getWine:local+'/api/wine/get',
    wineReleases:'https://api.github.com/repos/Gcenx/macOS_Wine_builds/releases',
    getDownloaded: local + '/api/downloaded',
    changeCurrent: local + '/api/wine/setcurrent',
    createBottle: local + '/api/wine/createBottle',
    runApp: local + '/api/wine/run',
}

const api = (type,props={}) => request(apiRoutes[type], props)

export default api