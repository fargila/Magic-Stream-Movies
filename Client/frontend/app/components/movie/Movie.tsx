
export const Movie = ({ movie })=> {
    return (
        <div className="col-md-4 mb-4">
            <div className="card h-100 shadow-sm">
                <div className="card-body d-flex flex flex-col">
                    <h5 className="card-title">{movie.title}</h5>
                    <p className="card-text mb-2">{movie.imdb_id}</p>
                </div>
            </div>
        </div>
    )
}