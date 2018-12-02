import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.List;

public class Day2Part1 {

	public static void main(String[] args) throws IOException {

		String inputPath = "";

		if (args.length != 0)
			inputPath = args[0];
		else
			inputPath = "input";

		List<String> ids = Files.readAllLines(Paths.get(inputPath));

		char[] alphabet = "abcdefghijklmnopqrstuvwxyz".toCharArray();

		int twoLettersRepetition = 0;
		int threeLettersRepetition = 0;

		boolean alreadyCountedTwo = false;
		boolean alreadyCountedThree = false;

		for (String id : ids) {
			char[] idChar = id.toCharArray();

			for (char letter : alphabet) {
				int occurences = 0;

				for (char idLetter : idChar) {
					if (idLetter == letter)
						occurences++;
				}

				if (occurences == 2 && !alreadyCountedTwo) {
					twoLettersRepetition++;
					alreadyCountedTwo = true;
				}

				if (occurences == 3 && !alreadyCountedThree) {
					threeLettersRepetition++;
					alreadyCountedThree = true;
				}

			}

			alreadyCountedTwo = false;
			alreadyCountedThree = false;

		}

		System.out.println(twoLettersRepetition * threeLettersRepetition);
	}
}
