package persona.backend;

import java.util.*;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

// RestContoller annotation means response + view?
@RestController
class PerfumeController {
	@GetMapping("/perfume")
	public static String handleRequest() {
		System.out.println("perfume handler got hit");
		return queryStock().toString();
	}

	private static List<String> queryStock() {
		List<String> collections = new ArrayList<>();
		collections.add("Maison Margiela Jazz Club");
		collections.add("Maison Margiela: By the Fireplace");
		collections.add("Prada Paradigme");
		collections.add("Club de nuit");
		collections.add("Lattafah Kamrah");
		collections.add("Bianco Latte");
		collections.add("Liquid Bum");
		return collections;
	}
}
