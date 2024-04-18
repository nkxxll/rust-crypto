use p256::{ecdh::diffie_hellman, SecretKey};
use rand::rngs::OsRng;

fn main() {
    let mut rng = OsRng;
    let priv_alice = SecretKey::random(&mut rng);
    let priv_bob = SecretKey::random(&mut rng);

    let pub_alice = priv_alice.public_key();
    let pub_bob = priv_bob.public_key();

    let secret = diffie_hellman(&priv_alice.to_nonzero_scalar(), pub_bob.as_affine());
    let secret_test = diffie_hellman(&priv_bob.to_nonzero_scalar(), pub_alice.as_affine());
    println!(
        "{:?} == {:?}",
        secret.raw_secret_bytes(),
        secret_test.raw_secret_bytes()
    );
}
