import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;

public class Day5Part2 {

	public static void main(String[] args) throws IOException {
		String inputPath = "";

		if (args.length != 0)
			inputPath = args[0];
		else
			inputPath = "input";

		byte[] input = Files.readAllBytes(Paths.get(inputPath));

		byte[] alphabet = "abcdefghijklmnopqrstuvwxyz".getBytes();

		int min = input.length;

		for (byte letter : alphabet) {
			int size = Polymer.react(Polymer.streamline(letter, input)).length;

			if (size < min)
				min = size;
		}

		System.out.println(min);

	}

}
