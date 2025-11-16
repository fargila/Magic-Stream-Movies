import Image from "next/image"

export const Movie = ({ movie })=> {
    return (
        <div className="col-md-4 mb-4">
            <div className="card h-100 shadow-sm">
                <div className="relative">
                    <Image 
                    className="contain-content h-[250px] w-full"
                    src={movie.poster_path} 
                    alt={movie.title} />
                </div>

                <div className="card-body d-flex flex flex-col">
                    <h5 className="card-title">{movie.title}</h5>
                    <p className="card-text mb-2">{movie.imdb_id}</p>
                </div>
                
                {movie.ranking?.ranking_name && (
                    <span className="bg-black m-3 p-2 text-white">
                        {movie.ranking.ranking_name}
                    </span>
                )}
            </div>
        </div>
    )
}