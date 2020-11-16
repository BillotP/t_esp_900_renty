<script lang="ts">
    export let offer;
</script>

<article class="media">
    <div class="media-left">
        <figure class="image is-64x64">
            <img
                src={(offer.assets.length && offer.assets[0].url) || 'https://bulma.io/images/placeholders/128x128.png'}
                alt="Rent offer view" />
        </figure>
    </div>
    <div class="media-content">
        <div class="content">
            <p>
                <strong>{offer.title[0].value}</strong>
                <small>{offer.price[0].value}</small>
                <small>{offer.price[0].currency}</small>
                <br />
                {(offer.description && offer.description[0].value) || '...'}
            </p>
        </div>
        <nav class="level is-mobile">
            <div class="level-left">
                <a href="/" class="level-item" aria-label="reply">
                    <span class="icon is-small">
                        <i class="fas fa-reply" aria-hidden="true" />
                    </span>
                </a>
                <a href="/" class="level-item" aria-label="retweet">
                    <span class="icon is-small">
                        <i class="fas fa-retweet" aria-hidden="true" />
                    </span>
                </a>
                <a href="/" class="level-item" aria-label="like">
                    <span class="icon is-small">
                        <i class="fas fa-heart" aria-hidden="true" />
                    </span>
                </a>
            </div>
        </nav>
    </div>
</article>
