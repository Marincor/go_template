async function teste() {
    const bla = await (await fetch('http://localhost:9090/health/check/list')).json()

    console.log(bla);
}
teste()